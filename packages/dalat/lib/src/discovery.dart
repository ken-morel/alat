import 'dart:async';
import 'dart:convert';
import 'dart:io'; // Import for RawDatagramSocket
import 'dart:ffi';

import 'package:ffi/ffi.dart';
import 'package:multicast_dns/multicast_dns.dart';

import 'bindings.dart';
import 'converters.dart'; // For Ip and Port types
import 'helpers.dart';
import 'models.dart'; // For FoundDevice and DeviceInfo

class DartDiscovery {
  final MDnsClient _client = MDnsClient(
    rawDatagramSocketFactory: (dynamic host, int port, {bool? reuseAddress, bool? reusePort, int? ttl}) {
      return RawDatagramSocket.bind(host, port, reuseAddress: true, reusePort: false, ttl: ttl ?? 255);
    },
  );
  Timer? _discoveryTimer;
  final int _instanceHandle; // To pass to Go FFI calls

  DartDiscovery(this._instanceHandle);

  void startDiscovery({Duration interval = const Duration(seconds: 2)}) {
    if (_discoveryTimer != null && _discoveryTimer!.isActive) {
      print('DartDiscovery already running.');
      return;
    }
    _discoveryTimer = Timer.periodic(interval, (timer) => _performDiscovery());
    _performDiscovery(); // Run immediately
  }

  void stopDiscovery() {
    _discoveryTimer?.cancel();
    _client.stop();
    print('DartDiscovery stopped.');
  }

  Future<void> _performDiscovery() async {
    print('Performing Dart mDNS discovery...');
    try {
      await _client.start();

      const String serviceName = '_alat._tcp.local';
      List<FoundDevice> foundDevices = [];

      await for (final PtrResourceRecord ptr
          in _client.lookup<PtrResourceRecord>(
            ResourceRecordQuery.serverPointer(serviceName),
          )) {
        await for (final SrvResourceRecord srv
            in _client.lookup<SrvResourceRecord>(
              ResourceRecordQuery.service(ptr.domainName),
            )) {
          await for (final IPAddressResourceRecord ipResource
              in _client.lookup<IPAddressResourceRecord>(
                ResourceRecordQuery.addressIPv4(srv.target),
              )) {
            final Ip ipAddress = ipResource.address.address;
            final Port port = srv.port;

            // Call Go FFI to get device info
            final ipJsonC = jsonEncode(ipAddress).toNativeUtf8();
            Pointer<Char> deviceInfoJsonC = nullptr;
            try {
              deviceInfoJsonC = bindings.query_device_info_json(
                ipJsonC.cast(),
                port,
              );
              if (deviceInfoJsonC != nullptr) {
                final String deviceInfoJson = deviceInfoJsonC
                    .cast<Utf8>()
                    .toDartString();
                final DeviceInfo deviceInfo = DeviceInfo.fromJson(
                  jsonDecode(deviceInfoJson),
                );
                bool found = false;
                for (final dev in foundDevices) {
                  if (dev.info.id == deviceInfo.id) {
                    found = true;
                    break;
                  }
                }
                if (!found) {
                  foundDevices.add(
                    FoundDevice(ip: ipAddress, port: port, info: deviceInfo),
                  );
                }
              } else {
                print(
                  'Error getting device info for $ipAddress:$port: ${InstanceHelpers.getAlatError()}',
                );
              }
            } finally {
              malloc.free(ipJsonC);
              if (deviceInfoJsonC != nullptr) {
                bindings.free_string(deviceInfoJsonC.cast());
              }
            }
          }
        }
      }

      // Report found devices to Go FFI
      if (foundDevices.isNotEmpty) {
        final devicesJson = jsonEncode(foundDevices);
        final devicesJsonC = devicesJson.toNativeUtf8();
        try {
          final result = bindings.discovery_provide_found_devices_json(
            _instanceHandle,
            devicesJsonC.cast(),
          );
          if (result != 0) {
            print(
              'Error reporting found devices to Go: $result, ${InstanceHelpers.getAlatError()}',
            );
          } else {
            print(
              'Successfully reported ${foundDevices.length} devices to Go.',
            );
          }
        } finally {
          malloc.free(devicesJsonC);
        }
      } else {
        print('No devices found.');
      }
    } catch (e) {
      print('Dart mDNS discovery error: $e');
    } finally {
      _client.stop(); // Stop client after each discovery cycle
    }
  }
}
