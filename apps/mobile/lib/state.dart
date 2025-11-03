import 'dart:async';
import 'dart:io';

import 'package:alat/services/transfer_notification_service.dart';
import 'package:dalat/dalat.dart' as dalat;
import 'package:flutter/material.dart';
import 'package:path_provider/path_provider.dart';
import 'package:file_picker/file_picker.dart';
import 'package:receive_sharing_intent/receive_sharing_intent.dart';

import 'services/notification_service.dart';

class PairRequestState {
  final dalat.PairRequest request;
  final Completer<dalat.PairResponse> completer;

  PairRequestState(this.request, this.completer);
}

class AppState extends ChangeNotifier {
  final NotificationService notificationService;
  final TransferNotificationService transferNotificationService;
  Timer? _transferStatusTimer;
  Timer? _webShareStatusTimer;

  dalat.AlatInstance? _alatInstance;
  dalat.AppConfig? _appSettings;
  dalat.ServiceConfig? _serviceSettings;
  dalat.WebshareStatus? _webShareStatus;

  final ValueNotifier<PairRequestState?> pendingPairRequest = ValueNotifier(
    null,
  );
  final ValueNotifier<List<String>> sharedFiles = ValueNotifier([]);

  AppState({
    required this.notificationService,
    required this.transferNotificationService,
  });

  dalat.AlatInstance? get node => _alatInstance;
  dalat.AppConfig? get settings => _appSettings;
  dalat.ServiceConfig? get serviceSettings => _serviceSettings;
  dalat.WebshareStatus? get webShareStatus => _webShareStatus;

  bool get isReady => _alatInstance != null && _appSettings != null;

  static Future<Directory> getAlatDir() async {
    try {
      return await getLibraryDirectory();
    } catch (e) {
      return await getApplicationSupportDirectory();
    }
  }

  static dalat.AppConfig createAppConfig() {
    final config = dalat.InstanceConfig.createAppConfig();
    config.deviceType = dalat.deviceTypeMobile;
    if (Platform.isLinux || Platform.isMacOS || Platform.isWindows) {
      config.deviceType = dalat.deviceTypeDesktop;
    }
    return config;
  }

  static Future<dalat.ServiceConfig> createServiceConfig() async {
    final config = dalat.InstanceConfig.createServiceConfig();
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
    } else {
      _alatInstance = dalat.AlatInstance.get(instances[0]);
    }

    _appSettings = await _alatInstance!.getAppConfig();
    if (_appSettings!.setupComplete) {
      _alatInstance!.start();
      _startTransferStatusUpdates();
      _startWebShareStatusUpdates();
      _alatInstance!.registerPairRequestHandler(pairRequestHandler);
    }
    _serviceSettings = await _alatInstance!.getServiceConfig();

    // Listen for incoming shared files
    ReceiveSharingIntent.instance.getMediaStream().listen((List<SharedMediaFile> value) {
      sharedFiles.value = value.map((f) => f.path).toList();
      notifyListeners();
    }, onError: (err) {
      print("getMediaStream error: $err");
    });

    ReceiveSharingIntent.instance.getInitialMedia().then((List<SharedMediaFile> value) {
      if (value.isNotEmpty) {
        sharedFiles.value = value.map((f) => f.path).toList();
        notifyListeners();
      }
    });

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
    _alatInstance!.registerPairRequestHandler(pairRequestHandler);
    _startTransferStatusUpdates();
    _startWebShareStatusUpdates();
    notifyListeners();
  }

  Future<void> saveSettings() async {
    await _alatInstance!.setAppConfig(_appSettings!);
    await _alatInstance!.setServiceConfig(_serviceSettings!);
  }

  void _startTransferStatusUpdates() {
    _transferStatusTimer?.cancel(); // Cancel any existing timer
    _transferStatusTimer = Timer.periodic(const Duration(seconds: 1), (
      timer,
    ) async {
      try {
        final status = await _alatInstance?.getFileTransfersStatus();
        if (status != null) {
          await transferNotificationService.showTransferProgress(status);
        }
      } catch (e) {
        // Handle or log the error appropriately
        print('Error fetching transfer status: $e');
      }
    });
  }

  void _startWebShareStatusUpdates() {
    _webShareStatusTimer?.cancel();
    _webShareStatusTimer = Timer.periodic(const Duration(seconds: 1), (
      timer,
    ) async {
      await _updateWebShareStatus();
    });
  }

  Future<void> _updateWebShareStatus() async {
    _webShareStatus = await _alatInstance?.getWebshareStatus();
    notifyListeners();
  }

  Future<void> webShareStart() async {
    await _alatInstance?.startWebshare();
    await _updateWebShareStatus();
  }

  Future<void> webShareStop() async {
    await _alatInstance?.stopWebshare();
    await _updateWebShareStatus();
  }

  Future<void> webShareSetPasscode(String passcode) async {
    await _alatInstance?.setWebsharePasscode(passcode);
    await _updateWebShareStatus();
  }

  Future<void> webShareAddFiles() async {
    FilePickerResult? result = await FilePicker.platform.pickFiles(
      allowMultiple: true,
    );
    if (result != null) {
      List<String> filePaths = result.files.map((file) => file.path!).toList();
      await _alatInstance?.addSharedFiles(filePaths);
      await _updateWebShareStatus();
    }
  }

  Future<void> webShareRemoveFile(String uuid) async {
    await _alatInstance?.removeSharedFilesByUUIDs([uuid]);
    await _updateWebShareStatus();
  }

  Future<void> webShareClearFiles() async {
    await _alatInstance?.clearSharedFiles();
    await _updateWebShareStatus();
  }

  @override
  void dispose() {
    _transferStatusTimer?.cancel();
    _webShareStatusTimer?.cancel();
    _alatInstance?.dispose();
    super.dispose();
  }

  Future<dalat.PairResponse> pairRequestHandler(dalat.PairRequest req) async {
    final completer = Completer<dalat.PairResponse>();

    notificationService.showPairingRequest(req);

    // Set the state that the UI will listen to.
    pendingPairRequest.value = PairRequestState(req, completer);

    return completer.future;
  }
}
