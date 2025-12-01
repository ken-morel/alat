import 'dart:async';
import 'dart:convert';
import 'dart:ffi';
import 'dart:isolate';

import 'package:dalat/src/bindings.dart';
import 'package:ffi/ffi.dart';

/// A request to be sent to the FFI-handling isolate.
class _FfiRequest {
  final int id;
  final String method;
  final List<dynamic> params;

  _FfiRequest(this.id, this.method, this.params);
}

/// A response received from the FFI-handling isolate.
class _FfiResponse {
  final int id;
  final dynamic result;
  final String? error;

  _FfiResponse(this.id, this.result, {this.error});
}

/// The entry point for the helper isolate.
///
/// This function listens for [_FfiRequest] messages, executes the corresponding
/// native function, and sends back a [_FfiResponse].
void _ffiIsolateEntry(SendPort mainSendPort) {
  final helperReceivePort = ReceivePort();
  mainSendPort.send(helperReceivePort.sendPort);

  helperReceivePort.listen((dynamic message) {
    if (message is _FfiRequest) {
      try {
        final result = _executeFfiCall(message.method, message.params);
        mainSendPort.send(_FfiResponse(message.id, result));
      } catch (e, s) {
        mainSendPort
            .send(_FfiResponse(message.id, null, error: '$e\n$s'));
      }
    }
  });
}

/// Executes the native function based on the method name and parameters.
///
/// This function acts as a router, converting string method names into
/// actual FFI calls.
dynamic _executeFfiCall(String method, List<dynamic> params) {
  // Helper to handle FFI calls that return a JSON string (Pointer<Char>).
  String? handleJsonStringResult(Pointer<Char> ptr) {
    if (ptr == nullptr) {
      return null;
    }
    try {
      return ptr.cast<Utf8>().toDartString();
    } finally {
      bindings.free_string(ptr);
    }
  }

  // Helper for simple FFI calls that return an integer status.
  void handleIntResult(int result) {
    if (result < 0) {
      throw Exception(
          'FFI call failed with code $result: ${bindings.get_error().cast<Utf8>().toDartString()}');
    }
  }

  switch (method) {
    // --- Instance Management ---
    case 'create_instance':
      final configPathC = (params[0] as String).toNativeUtf8();
      final appConfigC = (params[1] as String).toNativeUtf8();
      final serviceConfigC = (params[2] as String).toNativeUtf8();
      try {
        return bindings.create_instance(
            configPathC.cast(), appConfigC.cast(), serviceConfigC.cast());
      } finally {
        malloc.free(configPathC);
        malloc.free(appConfigC);
        malloc.free(serviceConfigC);
      }
    case 'start_instance':
      handleIntResult(bindings.start_instance(params[0]));
      return null;
    case 'stop_instance':
      bindings.stop_instance(params[0]);
      return null;
    case 'destroy_instance':
      bindings.destroy_instance(params[0]);
      return null;
    case 'get_instances':
      return handleJsonStringResult(bindings.get_instances());
    case 'get_node_status_json':
      return handleJsonStringResult(bindings.get_node_status_json(params[0]));

    // --- Config Management ---
    case 'get_app_config_json':
      return handleJsonStringResult(bindings.get_app_config_json(params[0]));
    case 'set_app_config_json':
      final jsonC = (params[1] as String).toNativeUtf8();
      try {
        handleIntResult(bindings.set_app_config_json(params[0], jsonC.cast()));
      } finally {
        malloc.free(jsonC);
      }
      return null;
    case 'get_service_config_json':
      return handleJsonStringResult(bindings.get_service_config_json(params[0]));
    case 'set_service_config_json':
      final jsonC = (params[1] as String).toNativeUtf8();
      try {
        handleIntResult(bindings.set_service_config_json(params[0], jsonC.cast()));
      } finally {
        malloc.free(jsonC);
      }
      return null;

    // --- Discovery ---
    case 'discovery_disable':
      handleIntResult(bindings.discovery_disable(params[0]));
      return null;
    case 'discovery_enabled':
      return bindings.discovery_enabled(params[0]);
    case 'discovery_provide_found_devices_json':
      final jsonC = (params[1] as String).toNativeUtf8();
      try {
        handleIntResult(bindings.discovery_provide_found_devices_json(params[0], jsonC.cast()));
      } finally {
        malloc.free(jsonC);
      }
      return null;

    // --- Pairing ---
    case 'request_pair_found_device':
      final deviceIdC = (params[1] as String).toNativeUtf8();
      try {
        return handleJsonStringResult(bindings.request_pair_found_device(params[0], deviceIdC.cast()));
      } finally {
        malloc.free(deviceIdC);
      }
    case 'submit_pair_response':
      final requestIdC = (params[1] as String).toNativeUtf8();
      final reasonC = (params[3] as String).toNativeUtf8();
      try {
        handleIntResult(bindings.submit_pair_response(params[0], requestIdC.cast(), params[2], reasonC.cast()));
      } finally {
        malloc.free(requestIdC);
        malloc.free(reasonC);
      }
      return null;
    case 'get_found_devices_json':
      return handleJsonStringResult(bindings.get_found_devices_json(params[0]));
    case 'get_paired_devices_json':
      return handleJsonStringResult(bindings.get_paired_devices_json(params[0]));
    case 'get_connected_devices_json':
      return handleJsonStringResult(bindings.get_connected_devices_json(params[0]));

    // --- File Transfer ---
    case 'query_connected_device_sysinfo':
      final deviceIdC = (params[1] as String).toNativeUtf8();
      try {
        return handleJsonStringResult(bindings.query_connected_device_sysinfo(params[0], deviceIdC.cast()));
      } finally {
        malloc.free(deviceIdC);
      }
    case 'query_send_files_to_connected_device':
      final deviceIdC = (params[1] as String).toNativeUtf8();
      final filesJsonC = (params[2] as String).toNativeUtf8();
      try {
        handleIntResult(bindings.query_send_files_to_connected_device(params[0], deviceIdC.cast(), filesJsonC.cast()));
      } finally {
        malloc.free(deviceIdC);
        malloc.free(filesJsonC);
      }
      return null;
    case 'get_file_transfers_status':
      return handleJsonStringResult(bindings.get_file_transfers_status(params[0]));

    // --- Webshare ---
    case 'wehsbare_get_status_json':
      return handleJsonStringResult(bindings.wehsbare_get_status_json(params[0]));
    case 'webshare_add_shared_files':
      final filesC = (params[1] as String).toNativeUtf8();
      try {
        handleIntResult(bindings.webshare_add_shared_files(params[0], filesC.cast()));
      } finally {
        malloc.free(filesC);
      }
      return null;
    case 'webshare_remove_shared_files_by_paths':
      final filesC = (params[1] as String).toNativeUtf8();
      try {
        handleIntResult(bindings.webshare_remove_shared_files_by_paths(params[0], filesC.cast()));
      } finally {
        malloc.free(filesC);
      }
      return null;
    case 'webshare_remove_shared_files_by_uuids':
      final filesC = (params[1] as String).toNativeUtf8();
      try {
        handleIntResult(bindings.webshare_remove_shared_files_by_uuids(params[0], filesC.cast()));
      } finally {
        malloc.free(filesC);
      }
      return null;
    case 'webshare_clear_shared_files':
      handleIntResult(bindings.webshare_clear_shared_files(params[0]));
      return null;
    case 'webshare_get_shared_files_json':
      return handleJsonStringResult(bindings.webshare_get_shared_files_json(params[0]));
    case 'webshare_get_passcode':
      return handleJsonStringResult(bindings.webshare_get_passcode(params[0]));
    case 'webshare_set_passcode':
      final passcodeC = (params[1] as String).toNativeUtf8();
      try {
        handleIntResult(bindings.webshare_set_passcode(params[0], passcodeC.cast()));
      } finally {
        malloc.free(passcodeC);
      }
      return null;
    case 'webshare_stop':
      bindings.webshare_stop(params[0]);
      return null;
    case 'webshare_start':
      bindings.webshare_start(params[0]);
      return null;

    // --- Other ---
    case 'get_alat_device_colors_json':
      return handleJsonStringResult(bindings.get_alat_device_colors_json());

    default:
      throw UnimplementedError('FFI method "$method" not implemented in isolate.');
  }
}

/// A class that manages the helper isolate and communication.
///
/// This provides a single `run` method to execute an FFI call on the helper
/// isolate and get the result back.
class FfiIsolate {
  static int _nextRequestId = 0;
  static final Map<int, Completer<_FfiResponse>> _requests = {};
  static Future<SendPort>? _sendPortFuture;

  static Future<SendPort> _getSendPort() {
    _sendPortFuture ??= () {
      final completer = Completer<SendPort>();
      final mainReceivePort = ReceivePort();

      mainReceivePort.listen((dynamic message) {
        if (message is SendPort) {
          completer.complete(message);
        } else if (message is _FfiResponse) {
          _requests[message.id]?.complete(message);
          _requests.remove(message.id);
        }
      });

      Isolate.spawn(_ffiIsolateEntry, mainReceivePort.sendPort);
      return completer.future;
    }();
    return _sendPortFuture!;
  }

  static Future<dynamic> run(String method, List<dynamic> params) async {
    final sendPort = await _getSendPort();
    final requestId = _nextRequestId++;
    final completer = Completer<_FfiResponse>();
    _requests[requestId] = completer;

    sendPort.send(_FfiRequest(requestId, method, params));

    final response = await completer.future;
    if (response.error != null) {
      throw Exception(
          'Error on FFI isolate for method $method: ${response.error}');
    }
    return response.result;
  }
}