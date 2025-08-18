import 'package:flutter/foundation.dart';
import 'package:alat/src/api/api_bridge.dart';

class DeviceProvider with ChangeNotifier {
  final ApiBridge _apiBridge = ApiBridge();
  List<dynamic> _devices = [];
  bool _isLoading = false;

  List<dynamic> get devices => _devices;
  bool get isLoading => _isLoading;

  Future<void> searchDevices() async {
    _isLoading = true;
    notifyListeners();

    try {
      _devices = await _apiBridge.searchDevices();
    } catch (e) {
      // Handle error
      print('Error searching for devices: $e');
      _devices = [];
    } finally {
      _isLoading = false;
      notifyListeners();
    }
  }
}
