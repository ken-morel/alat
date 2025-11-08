import 'dart:async';
import 'dart:convert';
import 'dart:ffi';

import 'package:ffi/ffi.dart';
import 'package:nsd/nsd.dart' as nsd;

import 'bindings.dart';
import 'converters.dart';
import 'helpers.dart';
import 'models.dart';

class DartDiscovery {
  final int _instanceHandle;
  StreamSubscription<nsd.Service>? _discoverySubscription;
  nsd.Discovery? _discovery;
  bool _isDiscovering = false;

  // Use a map to store found devices, keyed by "ip:port" to prevent duplicates.
  final Map<String, FoundDevice> _foundDevices = {};
  Timer? _reportingTimer;

  DartDiscovery(this._instanceHandle);

  Future<void> startDiscovery(
      {Duration reportingInterval = const Duration(seconds: 3)}) async {
    if (_isDiscovering) {
      return;
    }
    _isDiscovering = true;

    try {
      _discovery = await nsd.startDiscovery('_alat._tcp');
      _discoverySubscription =
          (_discovery! as Stream<nsd.Service>).listen(
        (service) {
          if (service.host != null && service.port != null) {
            // Service is already resolved
            _queryAndAddDevice(service.host!, service.port!);
          } else {
            // When a service is found, resolve it to get host and port
            nsd.resolve(service).then((resolvedService) {
              if (resolvedService.host != null && resolvedService.port != null) {
                _queryAndAddDevice(
                    resolvedService.host!, resolvedService.port!);
              }
            });
          }
        },
        onError: (e) {
          print('Discovery error: $e');
          // Consider stopping and restarting discovery on error
          stopDiscovery();
          startDiscovery(reportingInterval: reportingInterval);
        },
      );
    } catch (e) {
      print('Failed to start NSD discovery: $e');
    }

    // Periodically report the collected list of devices to the Go backend.
    _reportingTimer =
        Timer.periodic(reportingInterval, (_) => _reportDevicesToBackend());
  }

  void stopDiscovery() {
    if (!_isDiscovering) {
      return;
    }
    _isDiscovering = false;
    _discoverySubscription?.cancel();
    if (_discovery != null) {
      nsd.stopDiscovery(_discovery!);
    }
    _reportingTimer?.cancel();
    _foundDevices.clear();
  }

  Future<void> _queryAndAddDevice(Ip ipAddress, Port port) async {
    final key = '$ipAddress:$port';
    if (_foundDevices.containsKey(key)) {
      return; // Already found
    }

    final ipJsonC = jsonEncode(ipAddress).toNativeUtf8();
    Pointer<Char> deviceInfoJsonC = nullptr;
    try {
      deviceInfoJsonC = bindings.query_device_info_json(
        ipJsonC.cast(),
        port,
      );

      if (deviceInfoJsonC != nullptr) {
        final String deviceInfoJson =
            deviceInfoJsonC.cast<Utf8>().toDartString();
        final DeviceInfo deviceInfo =
            DeviceInfo.fromJson(jsonDecode(deviceInfoJson));

        _foundDevices[key] =
            FoundDevice(ip: ipAddress, port: port, info: deviceInfo);
        print('    -> Added device: ${deviceInfo.name} ($key)');
      } else {
        print(
            'Debug: Found unresponsive peer at $ipAddress:$port. Skipping. (Error: ${InstanceHelpers.getAlatError()})');
      }
    } finally {
      malloc.free(ipJsonC);
      if (deviceInfoJsonC != nullptr) {
        bindings.free_string(deviceInfoJsonC.cast());
      }
    }
  }

  void _reportDevicesToBackend() {
    final devices = _foundDevices.values.toList();
    if (devices.isEmpty) {
      return;
    }

    final devicesJson = jsonEncode(devices);
    final devicesJsonC = devicesJson.toNativeUtf8();
    try {
      final result = bindings.discovery_provide_found_devices_json(
        _instanceHandle,
        devicesJsonC.cast(),
      );
      if (result != 0) {
        print(
            'Error reporting found devices to Go: $result, ${InstanceHelpers.getAlatError()}');
      }
    } finally {
      malloc.free(devicesJsonC);
    }
  }
}
