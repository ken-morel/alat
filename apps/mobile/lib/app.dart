import 'package:alat/pages/dashboard/pair.dart';
import 'package:flutter/material.dart';
import 'package:alat/pages/dashboard/dashboard.dart';
import 'package:alat/pages/setup/setup.dart';
import 'package:alat/l10n/app_localizations.dart';
import 'package:alat/pages/start.dart';
import 'package:alat/theme.dart';
import 'package:provider/provider.dart';

class AlatApplication extends StatelessWidget {
  const AlatApplication({super.key});

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (_) => ThemeProvider(),
      child: Consumer<ThemeProvider>(
        builder: (context, themeProvider, child) {
          return MaterialApp(
            title: 'Alat',
            theme: createLightTheme(),
            darkTheme: createDarkTheme(),
            localizationsDelegates: AppLocalizations.localizationsDelegates,
            supportedLocales: AppLocalizations.supportedLocales,
            themeMode: themeProvider.themeMode,
            home: const StartPage(),
            routes: {
              '/dashboard': (context) => const DashboardPage(),
              '/dashboard/pair': (context) => const PairDevicePage(),
              '/setup': (context) => SetupAssistantPageView(),
            },
          );
        },
      ),
    );
  }
}
