import 'dart:async';
import 'dart:convert';
import 'dart:io';
import 'dart:ffi';

import 'package:ffi/ffi.dart';
import 'package:multicast_dns/multicast_dns.dart';

import 'bindings.dart';
import 'converters.dart';
import 'helpers.dart';
import 'models.dart';

class DartDiscovery {
  final MDnsClient _client = MDnsClient(
    rawDatagramSocketFactory: (dynamic host, int port,
        {bool? reuseAddress, bool? reusePort, int? ttl}) {
      // Allow multiple sockets to bind to the same port. This is necessary
      // on Linux to allow both discovery and advertising to coexist.
      return RawDatagramSocket.bind(host, port,
          reuseAddress: true, reusePort: true, ttl: ttl ?? 255);
    },
  );

  final int _instanceHandle;
  Timer? _reportingTimer;
  StreamSubscription? _discoverySubscription;
  bool _isDiscovering = false;

  // Use a map to store found devices, keyed by "ip:port" to prevent duplicates.
  final Map<String, FoundDevice> _foundDevices = {};

  DartDiscovery(this._instanceHandle);

  Future<void> startDiscovery(
      {Duration reportingInterval = const Duration(seconds: 3)}) async {
    if (_isDiscovering) {
      return;
    }
    _isDiscovering = true;
    // print('Starting Dart mDNS discovery...'); // Too noisy

    await _client.start();

    const String serviceType = '_alat._tcp.local';
    _discoverySubscription = _client
        .lookup<PtrResourceRecord>(ResourceRecordQuery.serverPointer(serviceType))
        .listen(
      (PtrResourceRecord ptr) {
        // Asynchronously resolve the service details.
        _resolveService(ptr.domainName);
      },
      onError: (dynamic error) {
        print('Dart mDNS discovery stream error: $error');
        // Consider restarting or handling the error.
      },
      onDone: () {
        // print('Dart mDNS discovery stream closed.'); // Too noisy
        if (_isDiscovering) {
          // If the stream closes unexpectedly, try to restart it.
          // This can happen on network changes.
          stopDiscovery();
          startDiscovery(reportingInterval: reportingInterval);
        }
      },
    );

    // Periodically report the collected list of devices to the Go backend.
    _reportingTimer =
        Timer.periodic(reportingInterval, (_) => _reportDevicesToBackend());
  }

  void stopDiscovery() {
    if (!_isDiscovering) {
      return;
    }
    // print('Stopping Dart mDNS discovery...'); // Too noisy
    _isDiscovering = false;
    _discoverySubscription?.cancel();
    _reportingTimer?.cancel();
    _client.stop();
    _foundDevices.clear();
  }

  Future<void> _resolveService(String serviceName) async {
    try {
      // Lookup SRV and A records for the given service name.
      await for (final SrvResourceRecord srv in _client
          .lookup<SrvResourceRecord>(ResourceRecordQuery.service(serviceName))) {
        await for (final IPAddressResourceRecord ip
            in _client.lookup<IPAddressResourceRecord>(
                ResourceRecordQuery.addressIPv4(srv.target))) {
          final deviceIp = ip.address.address;
          final devicePort = srv.port;
          final key = '$deviceIp:$devicePort';

          // If we already have this device, don't query it again.
          if (_foundDevices.containsKey(key)) {
            continue;
          }

          print('Discovered new device at $key. Querying info...');
          await _queryAndAddDevice(deviceIp, devicePort);
        }
      }
    } catch (e) {
      print('Error resolving service $serviceName: $e');
    }
  }

  Future<void> _queryAndAddDevice(Ip ipAddress, Port port) async {
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

        final key = '$ipAddress:$port';
        _foundDevices[key] =
            FoundDevice(ip: ipAddress, port: port, info: deviceInfo);
        print('    -> Added device: ${deviceInfo.name} ($key)');
      } else {
        print(
            'Error getting device info for $ipAddress:$port: ${InstanceHelpers.getAlatError()}');
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
      // print('No devices to report.');
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
      } else {
        print(
            'Successfully reported ${_foundDevices.length} devices to Go.');
      }
    } finally {
      malloc.free(devicesJsonC);
    }
  }
}
