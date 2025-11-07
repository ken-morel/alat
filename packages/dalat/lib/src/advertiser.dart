import 'package:nsd/nsd.dart';

class DartAdvertiser {
  Registration? _registration;

  Future<void> start(String deviceName, int port) async {
    if (_registration != null) {
      print('Advertiser already running.');
      return;
    }
    try {
      final service = Service(
        name: '${deviceName.replaceAll(RegExp(r'[^a-zA-Z0-9-]'), '')}-$port',
        type: '_alat._tcp',
        port: port,
      );
      _registration = await register(service);
      print('Service registered: $_registration');
    } catch (e) {
      print('Error registering service: $e');
      _registration = null;
    }
  }

  Future<void> stop() async {
    if (_registration != null) {
      try {
        await unregister(_registration!);
        print('Service unregistered: $_registration');
      } catch (e) {
        print('Error unregistering service: $e');
      } finally {
        _registration = null;
      }
    }
  }
}

