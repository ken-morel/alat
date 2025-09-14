import 'package:alat/state.dart';
import 'package:alat/pages/start.dart';
import 'package:alat/theme.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

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
      theme: lightTheme,
      darkTheme: darkTheme,
      themeMode: ThemeMode.system, // Or make this configurable
      home: const StartPage(),
      // TODO: Define routes for /dashboard and /setup
      routes: {
        // '/dashboard': (context) => const DashboardPage(),
        // '/setup': (context) => const SetupPage(),
      },
    );
  }
}

