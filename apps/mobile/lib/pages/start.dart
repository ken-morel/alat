import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'package:alat/state.dart';

class StartPage extends StatefulWidget {
  const StartPage({super.key});

  @override
  State<StartPage> createState() => _StartPageState();
}

class _StartPageState extends State<StartPage> {
  @override
  void initState() {
    super.initState();
    context.read<AppState>().addListener(_onAppStateChanged);
  }

  @override
  void dispose() {
    context.read<AppState>().removeListener(_onAppStateChanged);
    super.dispose();
  }

  void _onAppStateChanged() {
    final appState = context.read<AppState>();
    if (appState.isReady) {
      if (appState.settings!.setupComplete) {
      } else {
        Navigator.of(context).pushReplacementNamed('/setup');
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            CircularProgressIndicator(),
            SizedBox(height: 20),
            Text('Initializing Alat...'),
          ],
        ),
      ),
    );
  }
}
