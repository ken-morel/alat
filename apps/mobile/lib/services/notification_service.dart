import 'package:flutter_local_notifications/flutter_local_notifications.dart';

import 'navigation_service.dart';

class NotificationService {
  final FlutterLocalNotificationsPlugin _flutterLocalNotificationsPlugin =
      FlutterLocalNotificationsPlugin();
  final NavigationService _navigationService;

  NotificationService(this._navigationService);

  Future<void> init() async {
    const AndroidInitializationSettings initializationSettingsAndroid =
        AndroidInitializationSettings('@mipmap/ic_launcher');
    const LinuxInitializationSettings initializationSettingsLinux =
        LinuxInitializationSettings(defaultActionName: 'Open');

    const InitializationSettings initializationSettings = InitializationSettings(
      android: initializationSettingsAndroid,
      linux: initializationSettingsLinux,
    );

    await _flutterLocalNotificationsPlugin.initialize(
      initializationSettings,
      onDidReceiveNotificationResponse: onDidReceiveNotificationResponse,
    );
  }

  void onDidReceiveNotificationResponse(NotificationResponse response) {
    if (response.payload != null && response.payload!.isNotEmpty) {
      // The payload is the JSON of the PairRequest
      // We navigate to a route that knows how to handle this.
      // We will create this route in the next step.
      _navigationService.navigateTo('/pair-request', arguments: response.payload);
    }
  }

  Future<void> showPairingRequest(String deviceName, String payload) async {
    const AndroidNotificationDetails androidPlatformChannelSpecifics = AndroidNotificationDetails(
      'pairing_requests',
      'Pairing Requests',
      channelDescription: 'Notifications for new device pairing requests',
      importance: Importance.max,
      priority: Priority.high,
      showWhen: false,
    );
    const LinuxNotificationDetails linuxPlatformChannelSpecifics = LinuxNotificationDetails();

    const NotificationDetails platformChannelSpecifics = NotificationDetails(
      android: androidPlatformChannelSpecifics,
      linux: linuxPlatformChannelSpecifics,
    );

    await _flutterLocalNotificationsPlugin.show(
      0, // Notification ID
      'Alat Pairing Request',
      '$deviceName wants to connect.',
      platformChannelSpecifics,
      payload: payload,
    );
  }
}
