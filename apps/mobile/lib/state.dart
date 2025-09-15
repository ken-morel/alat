import 'dart:io';

import 'package:dalat/dalat.dart' as dalat;
import 'package:flutter/material.dart';
import 'package:path_provider/path_provider.dart';

class AppState extends ChangeNotifier {
  dalat.AlatInstance? _alatInstance;
  dalat.AppSettings? _appSettings;
  dalat.ServiceSettings? _serviceSettings;

  dalat.AlatInstance? get node => _alatInstance;
  dalat.AppSettings? get settings => _appSettings;
  dalat.ServiceSettings? get serviceSettings => _serviceSettings;

  bool get isReady => _alatInstance != null && _appSettings != null;

  Future<bool> initialize() async {
    final dir = await getApplicationDocumentsDirectory();
    final configPath = dir.path;
    final configDir = Directory(configPath);
    if (!await configDir.exists()) {
      await configDir.create(recursive: true);
    }

    _alatInstance = dalat.AlatInstance.create(
      configPath: configPath,
      deviceType: dalat.DeviceType.mobile,
    );

    _appSettings = await _alatInstance!.getAppSettings();
    if (_appSettings!.setupComplete) _alatInstance!.start();

    _serviceSettings = await _alatInstance!.getServiceSettings();
    _setupServices();

    // Notify listeners that initialization is complete.
    notifyListeners();
    return _appSettings!.setupComplete;
  }

  Future<void> _setupServices() async {
    final downloadsDir = await getDownloadsDirectory();
    if (downloadsDir != null) {
      _serviceSettings?.fileSend.saveFolder = downloadsDir.path;
    }
  }

  Future<void> completeSetup() async {
    if (_alatInstance == null) return;
    if (_appSettings == null) return;
    _appSettings!.setupComplete = true;
    await _alatInstance!.setAppSettings(_appSettings!);
    if (!_appSettings!.setupComplete) _alatInstance!.start();
    notifyListeners();
  }

  Future<void> saveSettings() async {
    await _alatInstance!.setAppSettings(_appSettings!);
    await _alatInstance!.setServiceSettings(_serviceSettings!);
  }

  @override
  void dispose() {
    _alatInstance?.dispose();
    super.dispose();
  }
}
