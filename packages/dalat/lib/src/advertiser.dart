import 'dart:io';
import 'package:multicast_dns/multicast_dns.dart';

class DartAdvertiser {
  final MDnsClient _client = MDnsClient(
    rawDatagramSocketFactory: (dynamic host, int port,
        {bool? reuseAddress, bool? reusePort, int? ttl}) {
      return RawDatagramSocket.bind(host, port,
          reuseAddress: true, reusePort: true, ttl: ttl ?? 255);
    },
  );
  bool _isAdvertising = false;

  Future<void> start(String deviceName, int port) async {
    if (_isAdvertising) {
      print('Advertiser already running.');
      return;
    }

    print('Starting advertiser using multicast_dns...');
    await _client.start();
    _isAdvertising = true;

    // Sanitize deviceName for use in the service instance name
    final String instanceName =
        '${deviceName.replaceAll(RegExp(r'[^a-zA-Z0-9-]'), '')}-$port';
    const String serviceType = '_alat._tcp';
    const String domain = 'local';

    final String fullInstanceName = '$instanceName.$serviceType.$domain';
    final String ptrQueryName = '$serviceType.$domain';

    // The hostname of the machine. This is the target for the SRV record.
    // Using Platform.localHostname and appending .local is a common convention.
    final String qualifiedHostName = '${Platform.localHostname}.$domain';

    print('Advertising service instance "$fullInstanceName" for host "$qualifiedHostName" on port $port');

    // Publish the PTR record to make the service type discoverable.
    // This points from _alat._tcp.local -> my-device-12345._alat._tcp.local
    _client.publish(PtrResourceRecord(ptrQueryName, domainName: fullInstanceName));

    // Publish the SRV record to point to the host and port.
    // This points from my-device-12345._alat._tcp.local -> my-linux-host.local:12345
    _client.publish(SrvResourceRecord(fullInstanceName,
        target: qualifiedHostName, port: port, priority: 0, weight: 0));

    // Publish A (IPv4) and AAAA (IPv6) records for our hostname so clients can find us.
    for (var interface in await NetworkInterface.list(
        includeLoopback: false, type: InternetAddressType.any)) {
      for (var addr in interface.addresses) {
        if (addr.type == InternetAddressType.ipv4) {
          _client.publish(IPAddressResourceRecord(qualifiedHostName, address: addr));
          print('  -> Publishing IPv4 address: ${addr.address}');
        }
        // Optionally publish IPv6 as well
        // if (addr.type == InternetAddressType.ipv6) {
        //   _client.publish(IPAddressResourceRecord(qualifiedHostName, address: addr));
        //   print('  -> Publishing IPv6 address: ${addr.address}');
        // }
      }
    }
  }

  void stop() {
    if (!_isAdvertising) {
      return;
    }
    print('Stopping advertiser...');
    _client.stop();
    _isAdvertising = false;
  }
}
