import 'dart:convert';
import 'dart:ffi'; // Still needed for FFI related types and utility functions
import 'dart:isolate';

import 'package:dalat/dalat.dart';
import 'package:dalat/src/helpers.dart';
import 'package:dalat/src/webshare.dart';
import 'package:ffi/ffi.dart';

import 'bindings.dart'; // Still needed as FfiIsolate.run will call bindings
import 'isolate_helper.dart'; // New import for FfiIsolate manager

/// A high-level, platform-agnostic API for interacting with the Alat core.
///
/// This class encapsulates the FFI handle management and provides a clean,
/// Dart-idiomatic interface to the underlying Go implementation.
import 'package:dalat/src/advertiser.dart';

class AlatInstance
    with InstanceHelpers, InstanceWebShare, InstancePair, InstanceConfig {
  @override
  final int handle;

  DartDiscovery? _dartDiscovery;
  DartAdvertiser? _dartAdvertiser;

  AlatInstance._(this.handle);

  static AlatInstance create({
    required String configPath,
    required ServiceConfig serviceConfig,
    required AppConfig appConfig,
  }) {
    final serviceString = jsonEncode(serviceConfig);
    final appString = jsonEncode(appConfig);
    final serviceC = serviceString.toNativeUtf8();
    final appC = appString.toNativeUtf8();
    final configC = configPath.toNativeUtf8();
    try {
      final handle = bindings.create_instance(
        configC.cast(),
        appC.cast(),
        serviceC.cast(),
      );
      if (handle <= 0) {
        throw Exception(
          'Failed to create handle, got id: $handle; ${InstanceHelpers.getAlatError()}',
        );
      }
      return AlatInstance._(handle);
    } finally {
      malloc.free(serviceC);
      malloc.free(appC);
      malloc.free(configC);
    }
  }

  static Future<AlatInstance> get(int handle) async {
    final instances = await AlatInstance.getInstances();
    if (instances.contains(handle)) {
      return AlatInstance._(handle);
    } else {
      throw "Instance $handle does not exist. in AlatInstance.get";
    }
  }

  static Future<List<int>> getInstances() async {
    final jsonStr = await FfiIsolate.run('get_instances', []) as String?;
    if (jsonStr == null || jsonStr.isEmpty) {
      return [];
    }
    return (jsonDecode(jsonStr) as List).map((k) => k as int).toList();
  }

  Future<void> start() async {
    // Note: discovery_disable is an FFI call, but it does not throw
    // if the module is disabled. We can call it freely.
    await helper('discovery_disable');

    // This call will throw on error via the isolate, and return null on success.
    await FfiIsolate.run('start_instance', [handle]);

    await helper('discovery_disable'); // Called again, as in original code

    final discoveryEnabled =
        await FfiIsolate.run('discovery_enabled', [handle]) as int? ?? -1;
    if (discoveryEnabled == 0) {
      startDartDiscovery();
      startAdvertising();
    }
    // The error check for start_instance is now handled by the await on FfiIsolate.run
  }

  Future<void> stop() async {
    stopDartDiscovery();
    stopAdvertising();
    await helper('stop_instance');
  }

  Future<void> dispose() async {
    _dartDiscovery?.stopDiscovery();
    _dartAdvertiser?.stop();
    await helper('destroy_instance');
    unregisterPairRequestHandler();
  }

  void startDartDiscovery({
    Duration reportingInterval = const Duration(seconds: 5),
  }) {
    _dartDiscovery ??= DartDiscovery(handle);
    _dartDiscovery!.startDiscovery(reportingInterval: reportingInterval);
  }

  void stopDartDiscovery() {
    _dartDiscovery?.stopDiscovery();
  }

  Future<void> startAdvertising() async {
    _dartAdvertiser ??= DartAdvertiser();
    final status = await getNodeStatus();
    final deviceName = (await getAppConfig()).deviceName;
    if (status.port > 0) {
      await _dartAdvertiser!.start(deviceName, status.port);
    } else {
      print('Could not start advertiser: port is not available.');
    }
  }

  void stopAdvertising() {
    _dartAdvertiser?.stop();
  }

  Future<NodeStatus> getNodeStatus() {
    return jsonGetterHelper('get_node_status_json', NodeStatus.fromJson);
  }

  Future<SysInfo> queryConnectedDeviceSysInfo(String deviceId) async {
    final infoJson = await FfiIsolate.run(
      'query_connected_device_sysinfo',
      [handle, deviceId],
    ) as String?;

    if (infoJson == null) {
      throw "Failed getting system information: ${InstanceHelpers.getAlatError()}";
    }
    return SysInfo.fromJson(jsonDecode(infoJson));
  }

  Future<void> querySendFilesToDevice(String deviceId, List<String> files) async {
    await FfiIsolate.run(
      'query_send_files_to_connected_device',
      [handle, deviceId, jsonEncode(files)],
    );
  }

  Future<FileTransfersStatus> getFileTransfersStatus() {
    return jsonGetterHelper(
      'get_file_transfers_status',
      FileTransfersStatus.fromJson,
    );
  }
}
