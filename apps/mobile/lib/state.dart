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

  static Future<Directory> getAlatDir() async {
    try {
      return await getLibraryDirectory();
    } catch (e) {
      return await getApplicationSupportDirectory();
    }
  }

  Future<bool> initialize() async {
    final instances = dalat.AlatInstance.getInstances();
    if (instances.isEmpty) {
      final configDir = await AppState.getAlatDir();
      if (!await configDir.exists()) {
        await configDir.create(recursive: true);
      }

      _alatInstance = dalat.AlatInstance.create(
        configPath: configDir.path,
        deviceType: dalat.deviceTypeMobile,
      );
    } else {
      _alatInstance = dalat.AlatInstance.get(instances[0]);
    }

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
    if (_alatInstance == null || _appSettings == null) {
      throw "Cannot complete setup, setings or instance not set";
    }
    _appSettings!.setupComplete = true;
    await _alatInstance!.setAppSettings(_appSettings!);
    _alatInstance!.start();
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
