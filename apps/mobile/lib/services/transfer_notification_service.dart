import 'package:flutter_local_notifications/flutter_local_notifications.dart';
import 'package:dalat/dalat.dart' as dalat;

class TransferNotificationService {
  final FlutterLocalNotificationsPlugin _flutterLocalNotificationsPlugin =
      FlutterLocalNotificationsPlugin();
  // Tracks ongoing progress notifications
  final Set<int> _activeProgressNotificationIds = {};
  // Tracks final notifications that have been shown to the user
  final Set<int> _shownFinalNotificationIds = {};

  Future<void> init() async {
    const AndroidInitializationSettings initializationSettingsAndroid =
        AndroidInitializationSettings('@mipmap/ic_launcher');
    const LinuxInitializationSettings initializationSettingsLinux =
        LinuxInitializationSettings(defaultActionName: 'Open');
    const InitializationSettings initializationSettings =
        InitializationSettings(
          android: initializationSettingsAndroid,
          linux: initializationSettingsLinux,
        );

    await _flutterLocalNotificationsPlugin.initialize(initializationSettings);
  }

  int _generateNotificationId(String deviceId, String fileName) {
    return (deviceId + fileName).hashCode & 0x7FFFFFFF;
  }

  Future<void> showTransferProgress(dalat.FileTransfersStatus status) async {
    final Set<int> currentProgressIds = {};

    // Combine sending and receiving into a single list for easier processing
    final allDeviceStatuses = [
      ...(status.sending ?? []).map((s) => {'status': s, 'type': 'Sending to'}),
      ...(status.receiving ?? []).map(
        (s) => {'status': s, 'type': 'Receiving from'},
      ),
    ];

    for (final item in allDeviceStatuses) {
      final deviceStatus = item['status'] as dalat.DeviceFileTransferStatus;
      final type = item['type'] as String;

      for (final transfer in deviceStatus.transfers ?? []) {
        final notificationId = _generateNotificationId(
          deviceStatus.device.id,
          transfer.fileName,
        );

        final isFinished =
            transfer.status == 'completed' || transfer.status == 'failed';

        if (isFinished) {
          // This is a final state.
          // 1. Cancel any lingering progress notification for this transfer.
          await _flutterLocalNotificationsPlugin.cancel(notificationId);
          _activeProgressNotificationIds.remove(notificationId);

          // 2. If we haven't shown the final notification yet, show it.
          if (!_shownFinalNotificationIds.contains(notificationId)) {
            final title = transfer.status == 'completed'
                ? 'Transfer complete'
                : 'Transfer failed';
            final body =
                '$type ${deviceStatus.device.name}: ${transfer.fileName}';
            await _showFinalNotification(notificationId, title, body);
            _shownFinalNotificationIds.add(notificationId);
          }
        } else {
          // This is an ongoing transfer.
          currentProgressIds.add(notificationId);
          await _showProgressNotification(
            notificationId,
            '$type ${deviceStatus.device.name}',
            transfer.fileName,
            transfer.percent.toInt(),
          );
        }
      }
    }

    // Clean up stale progress notifications for transfers that disappeared
    final notificationsToCancel = _activeProgressNotificationIds.difference(
      currentProgressIds,
    );
    for (final id in notificationsToCancel) {
      await _flutterLocalNotificationsPlugin.cancel(id);
    }

    _activeProgressNotificationIds.clear();
    _activeProgressNotificationIds.addAll(currentProgressIds);
  }

  Future<void> _showProgressNotification(
    int id,
    String title,
    String body,
    int progress,
  ) async {
    final AndroidNotificationDetails androidPlatformChannelSpecifics =
        AndroidNotificationDetails(
          'transfer_progress',
          'Transfer Progress',
          channelDescription: 'Notifications for file transfer progress',
          channelShowBadge: false,
          importance: Importance.low,
          priority: Priority.low,
          onlyAlertOnce: true,
          showProgress: true,
          maxProgress: 100,
          progress: progress,
        );

    const LinuxNotificationDetails linuxPlatformChannelSpecifics =
        LinuxNotificationDetails(urgency: LinuxNotificationUrgency.low);

    final NotificationDetails platformChannelSpecifics = NotificationDetails(
      android: androidPlatformChannelSpecifics,
      linux: linuxPlatformChannelSpecifics,
    );

    await _flutterLocalNotificationsPlugin.show(
      id,
      title,
      body,
      platformChannelSpecifics,
    );
  }

  Future<void> _showFinalNotification(int id, String title, String body) async {
    final AndroidNotificationDetails androidPlatformChannelSpecifics =
        AndroidNotificationDetails(
          'transfer_status',
          'Transfer Status',
          channelDescription: 'Notifications for completed or failed transfers',
          importance: Importance.defaultImportance,
          priority: Priority.defaultPriority,
        );

    const LinuxNotificationDetails linuxPlatformChannelSpecifics =
        LinuxNotificationDetails(urgency: LinuxNotificationUrgency.normal);

    final NotificationDetails platformChannelSpecifics = NotificationDetails(
      android: androidPlatformChannelSpecifics,
      linux: linuxPlatformChannelSpecifics,
    );

    await _flutterLocalNotificationsPlugin.show(
      id,
      title,
      body,
      platformChannelSpecifics,
    );
  }
}
