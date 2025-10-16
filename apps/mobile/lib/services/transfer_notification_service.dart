import 'package:flutter_local_notifications/flutter_local_notifications.dart';
import 'package:dalat/dalat.dart' as dalat;

class TransferNotificationService {
  final FlutterLocalNotificationsPlugin _flutterLocalNotificationsPlugin =
      FlutterLocalNotificationsPlugin();

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
    for (final deviceStatus in status.sending ?? []) {
      for (final transfer in deviceStatus.transfers ?? []) {
        final notificationId = _generateNotificationId(
          deviceStatus.device.id,
          transfer.fileName,
        );
        if (transfer.percent >= 100) {
          await _flutterLocalNotificationsPlugin.cancel(notificationId);
        } else {
          await _showProgressNotification(
            notificationId,
            'Sending to ${deviceStatus.device.name}',
            transfer.fileName,
            transfer.percent.toInt(),
          );
        }
      }
    }

    for (final deviceStatus in status.receiving ?? []) {
      for (final transfer in deviceStatus.transfers ?? []) {
        final notificationId = _generateNotificationId(
          deviceStatus.device.id,
          transfer.fileName,
        );
        if (transfer.percent >= 100) {
          await _flutterLocalNotificationsPlugin.cancel(notificationId);
        } else {
          await _showProgressNotification(
            notificationId,
            'Receiving from ${deviceStatus.device.name}',
            transfer.fileName,
            transfer.percent.toInt(),
          );
        }
      }
    }
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
}
