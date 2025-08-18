import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:alat/src/providers/device_provider.dart';
import 'package:alat/src/providers/theme_provider.dart';
import 'package:alat/src/screens/home_screen.dart';
import 'package:alat/src/theme/theme.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

void main() {
  runApp(
    MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (_) => ThemeProvider()),
        ChangeNotifierProvider(create: (_) => DeviceProvider()),
      ],
      child: const AlatApp(),
    ),
  );
}

class AlatApp extends StatelessWidget {
  const AlatApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Alat',
      theme: AppTheme.darkTheme,
      localizationsDelegates: AppLocalizations.localizationsDelegates,
      supportedLocales: AppLocalizations.supportedLocales,
      home: const HomeScreen(),
    );
  }
}
