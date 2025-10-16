import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:dalat/dalat.dart' as dalat;
import 'package:flutter/material.dart';
import 'package:path_provider/path_provider.dart';

import 'services/notification_service.dart';

/// A state object to hold a pending pair request and the completer to resolve it.
class PairRequestState {
  final dalat.PairRequest request;
  final Completer<dalat.PairResponse> completer;

  PairRequestState(this.request, this.completer);
}

class AppState extends ChangeNotifier {
  final NotificationService notificationService;

  dalat.AlatInstance? _alatInstance;
  dalat.AppConfig? _appSettings;
  dalat.ServiceConfig? _serviceSettings;

  final ValueNotifier<PairRequestState?> pendingPairRequest = ValueNotifier(
    null,
  );

  AppState({required this.notificationService});

  dalat.AlatInstance? get node => _alatInstance;
  dalat.AppConfig? get settings => _appSettings;
  dalat.ServiceConfig? get serviceSettings => _serviceSettings;

  bool get isReady => _alatInstance != null && _appSettings != null;

  static Future<Directory> getAlatDir() async {
    try {
      return await getLibraryDirectory();
    } catch (e) {
      return await getApplicationSupportDirectory();
    }
  }

  static dalat.AppConfig createAppConfig() {
    final config = dalat.AlatInstance.createAppConfig();
    config.deviceType = dalat.deviceTypeMobile;
    if (Platform.isLinux || Platform.isMacOS || Platform.isWindows) {
      config.deviceType = dalat.deviceTypeDesktop;
    }
    return config;
  }

  static Future<dalat.ServiceConfig> createServiceConfig() async {
    final config = dalat.AlatInstance.createServiceConfig();
    config.fileSend.saveFolder =
        ((await getDownloadsDirectory()) ?? Directory(".")).path;
    return config;
  }

  static Future<dalat.AlatInstance> createInstance(String configPath) async {
    final appConfig = AppState.createAppConfig();
    final serviceConfig = await AppState.createServiceConfig();
    return dalat.AlatInstance.create(
      configPath: configPath,
      serviceConfig: serviceConfig,
      appConfig: appConfig,
    );
  }

  Future<bool> initialize() async {
    final instances = dalat.AlatInstance.getInstances();
    if (instances.isEmpty) {
      final configDir = await AppState.getAlatDir();
      if (!await configDir.exists()) {
        await configDir.create(recursive: true);
      }
      _alatInstance = await AppState.createInstance(configDir.path);
      _alatInstance!.registerPairRequestHandler(pairRequestHandler);
    } else {
      _alatInstance = dalat.AlatInstance.get(instances[0]);
      _alatInstance!.registerPairRequestHandler(pairRequestHandler);
    }

    _appSettings = await _alatInstance!.getAppConfig();
    if (_appSettings!.setupComplete) {
      _alatInstance!.start();
    }
    _serviceSettings = await _alatInstance!.getServiceConfig();

    notifyListeners();
    return _appSettings!.setupComplete;
  }

  Future<void> completeSetup() async {
    if (_alatInstance == null || _appSettings == null) {
      throw "Cannot complete setup, setings or instance not set";
    }
    _appSettings!.setupComplete = true;
    await _alatInstance!.setAppConfig(_appSettings!);
    _alatInstance!.start();
    notifyListeners();
  }

  Future<void> saveSettings() async {
    await _alatInstance!.setAppConfig(_appSettings!);
    await _alatInstance!.setServiceConfig(_serviceSettings!);
  }

  @override
  void dispose() {
    _alatInstance?.dispose();
    super.dispose();
  }

  Future<dalat.PairResponse> pairRequestHandler(dalat.PairRequest req) async {
    final completer = Completer<dalat.PairResponse>();

    // The payload for the notification is the JSON representation of the request.
    final payload = jsonEncode(req);

    // Trigger the notification.
    notificationService.showPairingRequest(req.device.name, payload);

    // Set the state that the UI will listen to.
    pendingPairRequest.value = PairRequestState(req, completer);

    return completer.future;
  }
}
