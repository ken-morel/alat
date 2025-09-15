import 'package:alat/state.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'app.dart';

void main() async {
  // Ensure that plugin services are initialized so that `path_provider` works.
  WidgetsFlutterBinding.ensureInitialized();

  runApp(
    ChangeNotifierProvider(
      create: (context) => AppState(),
      child: const AlatApplication(),
    ),
  );
}
