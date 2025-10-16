import 'package:alat/services/notification_service.dart';
import 'package:alat/services/transfer_notification_service.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'app.dart';
import 'services/navigation_service.dart';
import 'state.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  // 1. Create services
  final navigationService = NavigationService();
  final notificationService = NotificationService(navigationService);
  await notificationService.init();
  final transferNotificationService = TransferNotificationService();
  await transferNotificationService.init();

  // 2. Create AppState, injecting the notification service
  final appState = AppState(
    notificationService: notificationService,
    transferNotificationService: transferNotificationService,
  );

  // 3. Initialize the Alat core
  await appState.initialize();

  runApp(
    ChangeNotifierProvider.value(
      value: appState,
      // 4. Pass the navigatorKey to the application widget
      child: AlatApplication(navigatorKey: navigationService.navigatorKey),
    ),
  );
}
