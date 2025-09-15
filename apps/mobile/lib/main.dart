import 'package:alat/pages/setup/setup.dart';
import 'package:alat/state.dart';
import 'package:alat/pages/start.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:alat/l10n/app_localizations.dart';

void main() async {
  // Ensure that plugin services are initialized so that `path_provider` works.
  WidgetsFlutterBinding.ensureInitialized();

  runApp(
    ChangeNotifierProvider(
      create: (context) => AppState()..initialize(),
      child: const MyApp(),
    ),
  );
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Alat',
      // theme: lightTheme,
      // darkTheme: darkTheme,
      localizationsDelegates: AppLocalizations.localizationsDelegates,
      supportedLocales: AppLocalizations.supportedLocales,
      themeMode: ThemeMode.system,
      home: const StartPage(),
      routes: {
        // '/dashboard': (context) => const DashboardPage(),
        '/setup': (context) => SetupAssistantPageView(),
      },
    );
  }
}
