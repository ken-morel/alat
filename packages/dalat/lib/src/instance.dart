import 'dart:convert';
import 'dart:ffi';

import 'package:ffi/ffi.dart';

import 'bindings.dart';
import 'models.dart';

/// A high-level, platform-agnostic API for interacting with the Alat core.
///
/// This class encapsulates the FFI handle management and provides a clean,
/// Dart-idiomatic interface to the underlying Go implementation.
class AlatInstance {
  final int _handle;

  AlatInstance._(this._handle);

  static AlatInstance create({
    required String configPath,
    required DeviceType deviceType,
  }) {
    final configPathC = configPath.toNativeUtf8();
    final handle = bindings.create_instance(
      configPathC.cast(),
      deviceType.index,
    );
    malloc.free(configPathC);

    if (handle == 0) {
      throw Exception('Failed to create AlatInstance in the Go core.');
    }
    return AlatInstance._(handle);
  }

  void start() {
    final result = bindings.start_instance(_handle);
    if (result != 0) {
      throw Exception('Failed to start AlatInstance. Code: $result');
    }
  }

  void stop() {
    bindings.stop_instance(_handle);
  }

  void dispose() {
    bindings.destroy_instance(_handle);
  }

  Future<AppSettings> getAppSettings() async {
    return _jsonHelper(bindings.get_app_settings_json, AppSettings.fromJson);
  }

  Future<void> setAppSettings(AppSettings settings) async {
    return _jsonSetterHelper(settings.toJson(), bindings.set_app_settings_json);
  }

  Future<ServiceSettings> getServiceSettings() async {
    return _jsonHelper(
      bindings.get_service_settings_json,
      ServiceSettings.fromJson,
    );
  }

  Future<void> setServiceSettings(ServiceSettings settings) async {
    return _jsonSetterHelper(
      settings.toJson(),
      bindings.set_service_settings_json,
    );
  }

  Future<List<FoundDevice>> getFoundDevices() async {
    return _jsonListHelper(
      bindings.get_found_devices_json,
      FoundDevice.fromJson,
    );
  }

  Future<List<PairedDevice>> getPairedDevices() async {
    return _jsonListHelper(
      bindings.get_paired_devices_json,
      PairedDevice.fromJson,
    );
  }

  Future<List<ConnectedDevice>> getConnectedDevices() async {
    return _jsonListHelper(
      bindings.get_connected_devices_json,
      ConnectedDevice.fromJson,
    );
  }

  Future<NodeStatus> getNodeStatus() async {
    return _jsonHelper(bindings.get_node_status_json, NodeStatus.fromJson);
  }

  // --- Private Helpers ---

  Future<T> _jsonHelper<T>(
    Pointer<Char> Function(int) ffiFunc,
    T Function(Map<String, dynamic>) fromJson,
  ) async {
    final ptr = ffiFunc(_handle);
    if (ptr == nullptr) {
      throw Exception('Failed to get data from Go core: function returned null pointer.');
    }
    try {
      final jsonStr = ptr.cast<Utf8>().toDartString();
      return fromJson(jsonDecode(jsonStr));
    } finally {
      bindings.free_string(ptr.cast());
    }
  }

  Future<List<T>> _jsonListHelper<T>(
    Pointer<Char> Function(int) ffiFunc,
    T Function(Map<String, dynamic>) fromJson,
  ) async {
    final ptr = ffiFunc(_handle);
    if (ptr == nullptr) {
      // An empty list is represented by a null pointer in this API
      return [];
    }
    try {
      final jsonStr = ptr.cast<Utf8>().toDartString();
      final List<dynamic> decoded = jsonDecode(jsonStr);
      return decoded.map((item) => fromJson(item as Map<String, dynamic>)).toList();
    } finally {
      bindings.free_string(ptr.cast());
    }
  }

  Future<void> _jsonSetterHelper(
    Map<String, dynamic> jsonData,
    int Function(int, Pointer<Char>) ffiFunc,
  ) async {
    final jsonStr = jsonEncode(jsonData);
    final jsonStrC = jsonStr.toNativeUtf8();
    try {
      final result = ffiFunc(_handle, jsonStrC.cast());
      if (result != 0) {
        throw Exception('Failed to set data in Go core. Code: $result');
      }
    } finally {
      malloc.free(jsonStrC);
    }
  }
}
