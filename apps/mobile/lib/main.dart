import 'package:alat/services/notification_service.dart';
import 'package:alat/services/transfer_notification_service.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'app.dart';
import 'services/navigation_service.dart';
import 'state.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  final navigationService = NavigationService();

  final notificationService = NotificationService(navigationService);
  await notificationService.init();
  final transferNotificationService = TransferNotificationService();
  await transferNotificationService.init();

  final appState = AppState(
    notificationService: notificationService,
    transferNotificationService: transferNotificationService,
  );
  try {
    await appState.initialize();
  } catch (e) {
    print("Error initializing application: $e");
  }

  runApp(
    ChangeNotifierProvider.value(
      value: appState,
      child: AlatApplication(navigatorKey: navigationService.navigatorKey),
    ),
  );
}
