import 'dart:convert';
import 'dart:ffi';
import 'dart:isolate';

import 'package:dalat/dalat.dart';
import 'package:ffi/ffi.dart';

import 'bindings.dart';
import 'pair.dart';

typedef PairRequestHandler = Future<PairResponse> Function(PairRequest);

/// A high-level, platform-agnostic API for interacting with the Alat core.
///
/// This class encapsulates the FFI handle management and provides a clean,
/// Dart-idiomatic interface to the underlying Go implementation.
class AlatInstance {
  final int handle;

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
          'Failed to create handle, got id: $handle; ${getAlatError()}',
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

  static String getAlatError() {
    final msgPointer = bindings.get_error();
    final error = msgPointer == nullptr
        ? "Unknown error"
        : msgPointer.cast<Utf8>().toDartString();
    bindings.free_string(msgPointer);
    return error;
  }

  void stop() {
    bindings.stop_instance(handle);
  }

  void registerPairRequestHandler(PairRequestHandler fun) {
    final nativeCallback =
        NativeCallable<
          Void Function(Int, Pointer<Char>, Pointer<Char>, Pointer<Char>)
        >.listener(_pairRequestHandler);
    bindings.register_async_pair_request_callback(
      handle,
      nativeCallback.nativeFunction,
    );
    _pairRequestHandlers[handle] = fun;
    _nativeCallables[handle] = nativeCallback;

    // Register the native callback with Go.
    bindings.register_async_pair_request_callback(
      handle,
      nativeCallback.nativeFunction,
    );
  }

  void dispose() {
    bindings.destroy_instance(handle);
    unregisterPairRequestHandler();
  }

  Future<AppConfig> getAppConfig() async {
    return _jsonHelper(bindings.get_app_config_json, AppConfig.fromJson);
  }

  Future<void> setAppConfig(AppConfig settings) async {
    return _jsonSetterHelper(settings.toJson(), bindings.set_app_config_json);
  }

  Future<ServiceConfig> getServiceConfig() async {
    return _jsonHelper(
      bindings.get_service_config_json,
      ServiceConfig.fromJson,
    );
  }

  Future<void> setServiceConfig(ServiceConfig settings) async {
    return _jsonSetterHelper(
      settings.toJson(),
      bindings.set_service_config_json,
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

  Future<List<DeviceColor>> getAlatColors() async {
    final ptr = bindings.get_alat_device_colors_json();
    if (ptr == nullptr) {
      return [];
    }
    try {
      final jsonStr = ptr.cast<Utf8>().toDartString();
      final List<dynamic> decoded = jsonDecode(jsonStr);
      return decoded
          .map((item) => DeviceColor.fromJson(item as Map<String, dynamic>))
          .toList();
    } finally {
      bindings.free_string(ptr.cast());
    }
  }

  // --- Private Helpers ---

  Future<T> _jsonHelper<T>(
    Pointer<Char> Function(int) ffiFunc,
    T Function(Map<String, dynamic>) fromJson,
  ) async {
    final ptr = ffiFunc(handle);
    if (ptr == nullptr) {
      throw Exception(
        'Failed to get data from Go core: function returned null pointer. ${getAlatError()}',
      );
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
    final ptr = ffiFunc(handle);
    if (ptr == nullptr) {
      // An empty list is represented by a null pointer in this API
      return [];
    }
    try {
      final jsonStr = ptr.cast<Utf8>().toDartString();
      final List<dynamic> decoded = jsonDecode(jsonStr) ?? [];
      return decoded
          .map((item) => fromJson(item as Map<String, dynamic>))
          .toList();
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
      final result = await Isolate.run(() => ffiFunc(handle, jsonStrC.cast()));
      if (result != 0) {
        throw Exception(
          'Failed to set data in Go core. Code: $result ${getAlatError()}',
        );
      }
    } finally {
      malloc.free(jsonStrC);
    }
  }

  Future<RequestPairResponse> requestPair(String deviceId) async {
    return await Isolate.run(() {
      final deviceIdC = deviceId.toNativeUtf8();
      final ptr = bindings.request_pair_found_device(handle, deviceIdC.cast());
      if (ptr == nullptr) {
        return RequestPairResponse(
          status: -1,
          error: "Alat sent no reponse",
          accepted: false,
          reason: "Could not query device",
        );
      } else {
        final result = ptr.cast<Utf8>().toDartString();
        return RequestPairResponse.fromJson(jsonDecode(result));
      }
    });
  }

  static AppConfig createAppConfig() {
    final ptr = bindings.default_app_config();
    if (ptr == nullptr) {
      throw "Could not create default app settings, backend sent invalid null response";
    }
    try {
      final jsonStr = ptr.cast<Utf8>().toDartString();
      final Map<String, dynamic> decoded = jsonDecode(jsonStr);
      return AppConfig.fromJson(decoded);
    } finally {
      bindings.free_string(ptr.cast());
    }
  }

  static ServiceConfig createServiceConfig() {
    final ptr = bindings.default_service_config();
    if (ptr == nullptr) {
      throw "Could not create default service configuration, backend sent invalid null response";
    }
    try {
      final jsonStr = ptr.cast<Utf8>().toDartString();
      final Map<String, dynamic> decoded = jsonDecode(jsonStr);
      return ServiceConfig.fromJson(decoded);
    } finally {
      bindings.free_string(ptr.cast());
    }
  }

  void unregisterPairRequestHandler() {
    _nativeCallables[handle]?.close();
    _nativeCallables.remove(handle);
    _pairRequestHandlers.remove(handle);
  }
}

final Map<int, PairRequestHandler> _pairRequestHandlers = {};
final Map<int, NativeCallable> _nativeCallables = {};
void _pairRequestHandler(
  int handle,
  Pointer<Char> requestIdC,
  Pointer<Char> pairTokenC,
  Pointer<Char> deviceDetailsC,
) {
  final handler = _pairRequestHandlers[handle];
  try {
    if (handler == null) return;
    final requestId = requestIdC.cast<Utf8>().toDartString();
    final pairToken = Uint8ListConverter().fromJson(
      jsonDecode(pairTokenC.cast<Utf8>().toDartString()),
    );
    final deviceDetails = DeviceDetails.fromJson(
      jsonDecode(deviceDetailsC.cast<Utf8>().toDartString()),
    );

    // The handler is already being called on the correct isolate, so
    // another `Isolate.run` is not necessary.
    handler(
      PairRequest(
        requestid: requestId,
        token: pairToken,
        device: deviceDetails,
      ),
    ).then((response) {
      final newRequestIdC = requestId.toNativeUtf8();
      final reasonC = response.reason.toNativeUtf8();
      try {
        bindings.submit_pair_response(
          handle,
          newRequestIdC.cast(),
          response.accepted,
          reasonC.cast(),
        );
      } finally {
        malloc.free(newRequestIdC);
        malloc.free(reasonC);
      }
    });
  } finally {
    malloc.free(requestIdC);
    malloc.free(pairTokenC);
    malloc.free(deviceDetailsC);
  }
}
