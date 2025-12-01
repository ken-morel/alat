import 'dart:convert';
import 'dart:ffi';

import 'package:dalat/dalat.dart';
import 'package:dalat/src/bindings.dart';
import 'package:dalat/src/helpers.dart';
import 'package:dalat/src/helpers.dart';
import 'package:dalat/src/isolate_helper.dart'; // New import
import 'package:ffi/ffi.dart';
import 'package:json_annotation/json_annotation.dart';
part 'pair.g.dart';

typedef PairRequestHandler = Future<PairResponse> Function(PairRequest);

mixin InstancePair on InstanceHelpers {
  void registerPairRequestHandler(PairRequestHandler fun) {
    final nativeCallback =
        NativeCallable<
          Void Function(Int, Pointer<Char>, Pointer<Char>, Pointer<Char>)
        >.listener(pairRequestHandler);
    bindings.register_async_pair_request_callback(
      handle,
      nativeCallback.nativeFunction,
    );
    pairRequestHandlers[handle] = fun;
    nativeCallables[handle] = nativeCallback;

    // Register the native callback with Go. (This line is duplicated in original, keeping for now)
    bindings.register_async_pair_request_callback(
      handle,
      nativeCallback.nativeFunction,
    );
  }

  void unregisterPairRequestHandler() {
    nativeCallables[handle]?.close();
    nativeCallables.remove(handle);
    pairRequestHandlers.remove(handle);
  }

  Future<RequestPairResponse> requestPair(String deviceId) async {
    final resultJson = await FfiIsolate.run(
      'request_pair_found_device',
      [handle, deviceId],
    ) as String?;

    if (resultJson == null) {
      return RequestPairResponse(
        status: -1,
        error: "Alat sent no response",
        accepted: false,
        reason: "Could not query device",
      );
    } else {
      return RequestPairResponse.fromJson(jsonDecode(resultJson));
    }
  }

  Future<List<FoundDevice>> getFoundDevices() {
    return jsonListGetterHelper(
      'get_found_devices_json',
      FoundDevice.fromJson,
    );
  }

  Future<List<PairedDevice>> getPairedDevices() {
    return jsonListGetterHelper(
      'get_paired_devices_json',
      PairedDevice.fromJson,
    );
  }

  Future<List<ConnectedDevice>> getConnectedDevices() {
    return jsonListGetterHelper(
      'get_connected_devices_json',
      ConnectedDevice.fromJson,
    );
  }
}

@JsonSerializable()
class PairResponse {
  @JsonKey(name: 'accepted')
  bool accepted;
  @JsonKey(name: 'reason')
  String reason;
  PairResponse({required this.accepted, required this.reason});
  Map<String, dynamic> toJson() => _$PairResponseToJson(this);
  factory PairResponse.fromJson(Map<String, dynamic> json) =>
      _$PairResponseFromJson(json);
}

@JsonSerializable()
class PairRequest {
  @JsonKey(name: 'requestId')
  String requestid;
  @JsonKey(name: 'token')
  @Uint8ListConverter()
  PairToken token;
  @JsonKey(name: 'device')
  DeviceDetails device;
  PairRequest({
    required this.requestid,
    required this.token,
    required this.device,
  });
  Map<String, dynamic> toJson() => _$PairRequestToJson(this);
  factory PairRequest.fromJson(Map<String, dynamic> json) =>
      _$PairRequestFromJson(json);
}

final Map<int, PairRequestHandler> pairRequestHandlers = {};
final Map<int, NativeCallable> nativeCallables = {};
void pairRequestHandler(
  int handle,
  Pointer<Char> requestIdC,
  Pointer<Char> pairTokenC,
  Pointer<Char> deviceDetailsC,
) {
  final handler = pairRequestHandlers[handle];
  try {
    if (handler == null) return;
    final requestId = requestIdC.cast<Utf8>().toDartString();
    final pairToken = Uint8ListConverter().fromJson(
      jsonDecode(pairTokenC.cast<Utf8>().toDartString()),
    );
    final deviceDetails = DeviceDetails.fromJson(
      jsonDecode(deviceDetailsC.cast<Utf8>().toDartString()),
    );

    handler(
      PairRequest(
        requestid: requestId,
        token: pairToken,
        device: deviceDetails,
      ),
    ).then((response) {
      // This FFI call now goes through the helper isolate
      FfiIsolate.run(
        'submit_pair_response',
        [handle, requestId, response.accepted, response.reason],
      ).catchError((e) {
        print('Error submitting pair response via isolate: $e');
      });
    });
  } finally {
    // These are freed by the FFI callback mechanism, not by us.
    // malloc.free(requestIdC);
    // malloc.free(pairTokenC);
    // malloc.free(deviceDetailsC);
  }
}
