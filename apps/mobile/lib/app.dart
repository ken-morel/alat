import 'package:flutter/material.dart';
import 'package:alat/pages/dashboard/dashboard.dart';
import 'package:alat/pages/setup/setup.dart';
import 'package:alat/l10n/app_localizations.dart';
import 'package:alat/pages/start.dart';

class AlatApplication extends StatelessWidget {
  const AlatApplication({super.key});

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
        '/dashboard': (context) => const DashboardPage(),
        '/setup': (context) => SetupAssistantPageView(),
      },
    );
  }
}
