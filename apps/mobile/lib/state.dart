import 'dart:io';

import 'package:dalat/dalat.dart';
import 'package:flutter/material.dart';
import 'package:path_provider/path_provider.dart';

class AppState extends ChangeNotifier {
  AlatInstance? _alatInstance;
  AppSettings? _appSettings;

  AlatInstance? get alatInstance => _alatInstance;
  AppSettings? get appSettings => _appSettings;

  bool get isReady => _alatInstance != null && _appSettings != null;

  Future<void> initialize() async {
    final dir = await getApplicationDocumentsDirectory();
    final configPath = dir.path;

    // Ensure the config directory exists
    final configDir = Directory(configPath);
    if (!await configDir.exists()) {
      await configDir.create(recursive: true);
    }

    _alatInstance = AlatInstance.create(
      configPath: configPath,
      deviceType: DeviceType.mobile,
    );

    _appSettings = await _alatInstance!.getAppSettings();

    if (_appSettings!.setupComplete) {
      _alatInstance!.start();
    }

    // Notify listeners that initialization is complete.
    notifyListeners();
  }

  Future<void> completeSetup(AppSettings newSettings) async {
    if (_alatInstance == null) return;

    await _alatInstance!.setAppSettings(newSettings);
    _appSettings = newSettings;

    // Start the node after setup is complete
    _alatInstance!.start();

    notifyListeners();
  }

  @override
  void dispose() {
    _alatInstance?.dispose();
    super.dispose();
  }
}
