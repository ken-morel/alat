import 'dart:convert';
import 'dart:ffi';
import 'dart:isolate';

import 'package:dalat/dalat.dart';
import 'package:dalat/src/helpers.dart';
import 'package:dalat/src/webshare.dart';
import 'package:ffi/ffi.dart';

import 'bindings.dart';

/// A high-level, platform-agnostic API for interacting with the Alat core.
///
/// This class encapsulates the FFI handle management and provides a clean,
/// Dart-idiomatic interface to the underlying Go implementation.
class AlatInstance
    with InstanceHelpers, InstanceWebShare, InstancePair, InstanceConfig {
  @override
  final int handle;

  DartDiscovery? _dartDiscovery;

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

  factory AlatInstance.get(int handle) {
    final instances = AlatInstance.getInstances();
    if (instances.contains(handle)) {
      return AlatInstance._(handle);
    } else {
      throw "Instance $handle does not exist. in AlatInstance.get";
    }
  }

  static List<int> getInstances() {
    final ptr = bindings.get_instances();
    if (ptr == nullptr) {
      return [];
    }
    try {
      final jsonStr = ptr.cast<Utf8>().toDartString();
      return (jsonDecode(jsonStr) as List).map((k) => k as int).toList();
    } finally {
      bindings.free_string(ptr.cast());
    }
  }

  void start() {
    final result = bindings.start_instance(handle);
    if (bindings.discovery_enabled(handle) == 0) {
      startDartDiscovery();
    }
    if (result != 0) {
      final msgPointer = bindings.get_error();
      try {
        final error = msgPointer == nullptr
            ? "Unknown error"
            : msgPointer.cast<Utf8>().toDartString();
        throw Exception(
          'Failed to start AlatInstance. Code: $result, Error: $error',
        );
      } finally {
        bindings.free_string(msgPointer);
      }
    }
  }

  void stop() {
    stopDartDiscovery();
    bindings.stop_instance(handle);
  }

  void dispose() {
    _dartDiscovery?.stopDiscovery();
    bindings.destroy_instance(handle);
    unregisterPairRequestHandler();
  }

  void startDartDiscovery({Duration interval = const Duration(seconds: 5)}) {
    _dartDiscovery ??= DartDiscovery(handle);
    _dartDiscovery!.startDiscovery(interval: interval);
  }

  void stopDartDiscovery() {
    _dartDiscovery?.stopDiscovery();
  }

  Future<NodeStatus> getNodeStatus() {
    return jsonGetterHelper(bindings.get_node_status_json, NodeStatus.fromJson);
  }

  Future<SysInfo> queryConnectedDeviceSysInfo(String deviceId) {
    return Isolate.run(() {
      final deviceIdC = deviceId.toNativeUtf8();
      try {
        final infoC = bindings.query_connected_device_sysinfo(
          handle,
          deviceIdC.cast(),
        );
        if (infoC == nullptr) {
          throw "Failed getting system information: ${InstanceHelpers.getAlatError()}";
        }
        try {
          final info = infoC.cast<Utf8>().toDartString();
          return SysInfo.fromJson(jsonDecode(info));
        } finally {
          bindings.free_string(infoC);
        }
      } finally {
        malloc.free(deviceIdC);
      }
    });
  }

  Future<void> querySendFilesToDevice(String deviceId, List<String> files) {
    return Isolate.run(() {
      final deviceIdC = deviceId.toNativeUtf8();
      final filesJsonC = jsonEncode(files).toNativeUtf8();
      try {
        final response = bindings.query_send_files_to_connected_device(
          handle,
          deviceIdC.cast(),
          filesJsonC.cast(),
        );
        if (response < 0) {
          throw "Error sending files: ${InstanceHelpers.getAlatError()}";
        }
      } finally {
        malloc.free(deviceIdC);
        malloc.free(filesJsonC);
      }
    });
  }

  Future<FileTransfersStatus> getFileTransfersStatus() {
    return jsonGetterHelper(
      bindings.get_file_transfers_status,
      FileTransfersStatus.fromJson,
    );
  }
}
