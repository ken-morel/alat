import 'package:alat/services/notification_service.dart';
import 'package:alat/services/transfer_notification_service.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'app.dart';
import 'services/navigation_service.dart';
import 'state.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  final notificationService = NotificationService(navigatorKey);
  await notificationService.init();

  final transferNotificationService = TransferNotificationService();
  await transferNotificationService.init();

  runApp(
    ChangeNotifierProvider(
      create: (context) => AppState(
        notificationService: notificationService,
        transferNotificationService: transferNotificationService,
      ),
      child: AlatApplication(navigatorKey: navigatorKey),
    ),
  );
}
