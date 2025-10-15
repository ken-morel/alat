import 'dart:async';

import 'package:alat/state.dart';
import 'package:flutter/material.dart';
import 'package:dalat/dalat.dart' as dalat;
import 'package:provider/provider.dart';

class FoundDevicesList extends StatefulWidget {
  final void Function(dalat.FoundDevice) onConnectionUserRequest;
  const FoundDevicesList({super.key, required this.onConnectionUserRequest});

  @override
  State<FoundDevicesList> createState() => _FoundDeviceListState();
}

class _FoundDeviceListState extends State<FoundDevicesList> {
  List<dalat.FoundDevice> _foundDevices = [];
  Timer? _timer;
  late AppState _appState;

  @override
  void initState() {
    super.initState();
    _appState = context.read<AppState>();
    _startDeviceDiscovery();
  }

  @override
  void dispose() {
    _timer?.cancel();
    super.dispose();
  }

  void _startDeviceDiscovery() {
    _fetchDevices();
    _timer = Timer.periodic(const Duration(seconds: 1), (_) => _fetchDevices());
  }

  Future<void> _fetchDevices() async {
    final devices = await _appState.node?.getFoundDevices();
    if (devices != null && mounted) {
      setState(() {
        _foundDevices = devices;
      });
    }
  }

  IconData _deviceTypeToIcon(dalat.DeviceType type) {
    switch (type) {
      case "desktop":
        return Icons.desktop_windows;
      case "mobile":
        return Icons.phone_iphone;
      default:
        return Icons.devices_other;
    }
  }

  void _showDeviceDetailsSheet(BuildContext context, dalat.FoundDevice device) {
    showModalBottomSheet(
      context: context,
      isScrollControlled: true,
      builder: (BuildContext context) {
        return Padding(
          padding: const EdgeInsets.all(24.0),
          child: Wrap(
            children: [
              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Row(
                    children: [
                      Icon(
                        _deviceTypeToIcon(device.info.type),
                        size: 32,
                        color: Theme.of(context).colorScheme.primary,
                      ),
                      const SizedBox(width: 16),
                      Expanded(
                        child: Text(
                          device.info.name,
                          style: Theme.of(context).textTheme.headlineSmall,
                        ),
                      ),
                      IconButton(
                        icon: const Icon(Icons.close),
                        onPressed: () => Navigator.pop(context),
                      ),
                    ],
                  ),
                  const SizedBox(height: 24),
                  Text(
                    'Details',
                    style: Theme.of(context).textTheme.labelLarge,
                  ),
                  const Divider(),
                  ListTile(
                    title: const Text('Color'),
                    subtitle: Text(device.info.color.name),
                  ),
                  ListTile(
                    title: const Text('ID'),
                    subtitle: Text(
                      device.info.id,
                      style: const TextStyle(fontFamily: 'monospace'),
                    ),
                  ),
                  ListTile(
                    title: const Text('IP Address'),
                    subtitle: Text(device.ip),
                  ),
                  ListTile(
                    title: const Text('Port'),
                    subtitle: Text(device.port.toString()),
                  ),
                  const SizedBox(height: 24),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.end,
                    children: [
                      OutlinedButton(
                        onPressed: () => Navigator.pop(context),
                        child: const Text('Cancel'),
                      ),
                      const SizedBox(width: 12),
                      ElevatedButton(
                        onPressed: () {
                          Navigator.pop(context);
                          widget.onConnectionUserRequest(device);
                        },
                        child: const Text('Connect'),
                      ),
                    ],
                  ),
                ],
              ),
            ],
          ),
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return _foundDevices.isEmpty
        ? _buildNoDevicesFound(context)
        : _buildFoundDevicesList(context);
  }

  Widget _buildNoDevicesFound(BuildContext context) {
    return Center(
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Icon(Icons.search, size: 80, color: Colors.grey.shade400),
            const SizedBox(height: 20),
            Text(
              "Searching for devices...",
              style: Theme.of(context).textTheme.headlineSmall,
              textAlign: TextAlign.center,
            ),
            const SizedBox(height: 10),
            Text(
              "Make sure other devices are on the same Wi-Fi network and have Alat open.",
              textAlign: TextAlign.center,
              style: Theme.of(context).textTheme.bodyMedium,
            ),
            const SizedBox(height: 30),
            const CircularProgressIndicator(),
          ],
        ),
      ),
    );
  }

  Widget _buildFoundDevicesList(BuildContext context) {
    return ListView.builder(
      itemCount: _foundDevices.length,
      itemBuilder: (BuildContext context, int index) {
        final device = _foundDevices[index];
        final deviceColor = Color.fromRGBO(
          device.info.color.r,
          device.info.color.g,
          device.info.color.b,
          1,
        );
        return Card(
          margin: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
          child: ListTile(
            leading: CircleAvatar(backgroundColor: deviceColor),
            title: Text(device.info.name),
            subtitle: Text(device.info.type),
            onTap: () => _showDeviceDetailsSheet(context, device),
          ),
        );
      },
    );
  }
}
