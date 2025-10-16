import 'dart:convert';

import 'package:flutter_local_notifications/flutter_local_notifications.dart';
import 'package:dalat/dalat.dart' as dalat;

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

    const InitializationSettings initializationSettings =
        InitializationSettings(
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
      _navigationService.navigateTo(
        '/pair-request',
        arguments: response.payload,
      );
    }
  }

  Future<void> showPairingRequest(dalat.PairRequest req) async {
    final payload = jsonEncode(req);
    const AndroidNotificationDetails androidPlatformChannelSpecifics =
        AndroidNotificationDetails(
          'pairing_requests',
          'Pairing Requests',
          channelDescription: 'Notifications for new device pairing requests',
          importance: Importance.max,
          priority: Priority.high,
          showWhen: false,
        );

    const LinuxNotificationDetails linuxPlatformChannelSpecifics =
        LinuxNotificationDetails(
          urgency: LinuxNotificationUrgency.critical, // Make it more prominent
        );

    const NotificationDetails platformChannelSpecifics = NotificationDetails(
      android: androidPlatformChannelSpecifics,
      linux: linuxPlatformChannelSpecifics,
    );

    await _flutterLocalNotificationsPlugin.show(
      0, // Notification ID
      'Alat Pairing Request',
      '${req.device.name} of color ${req.device.color.name} wants to connect',
      platformChannelSpecifics,
      payload: payload,
    );
  }
}
