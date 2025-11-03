import 'dart:async';

import 'package:alat/components/deviceicon.dart';
import 'package:alat/state.dart';
import 'package:flutter/material.dart';
import 'package:dalat/dalat.dart' as dalat;
import 'package:provider/provider.dart';

class ConnecedDevicesSelection extends StatefulWidget {
  final void Function(List<dalat.ConnectedDevice>) onChange;
  final List<dalat.ConnectedDevice> selecedDevices;
  const ConnecedDevicesSelection({
    super.key,
    required this.onChange,
    this.selecedDevices = const [],
  });

  @override
  State<ConnecedDevicesSelection> createState() =>
      _ConnectedDeviceSelectionState();
}

class _ConnectedDeviceSelectionState extends State<ConnecedDevicesSelection> {
  late AppState _appState;
  late List<dalat.ConnectedDevice> connectedDevices;
  late List<dalat.ConnectedDevice> choosedDevices;
  Timer? _timer;

  @override
  void initState() {
    _appState = context.read<AppState>();
    connectedDevices = widget.selecedDevices;
    choosedDevices = connectedDevices;
    _timer = Timer.periodic(const Duration(seconds: 1), (_) => _fetchDevices());
    _fetchDevices();
    super.initState();
  }

  @override
  void dispose() {
    _timer?.cancel();
    super.dispose();
  }

  Future<void> _fetchDevices() async {
    final devices = await _appState.node?.getConnectedDevices();
    if (devices != null && mounted) {
      setState(() {
        connectedDevices = devices;
        connectedDevices.sort((a, b) {
          var list = [a.info.id, b.info.id];
          list.sort();
          return list[0] == a.info.id ? 1 : -1;
        });
      });
    }
  }

  void _showDeviceDetailsSheet(
    BuildContext context,
    dalat.ConnectedDevice device,
  ) {
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
                      DeviceIcon(
                        deviceType: device.info.type,
                        color: device.info.color,
                        size: 32,
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
    return connectedDevices.isEmpty
        ? _buildNoDevicesFound(context)
        : _buildConnecedDevicesSelection(context);
  }

  Widget _buildNoDevicesFound(BuildContext context) {
    return Center(
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              "No connected device",
              style: Theme.of(
                context,
              ).textTheme.headlineSmall?.copyWith(color: Colors.grey),
              textAlign: TextAlign.center,
            ),
            const SizedBox(height: 30),
            const CircularProgressIndicator(),
          ],
        ),
      ),
    );
  }

  Widget _buildConnecedDevicesSelection(BuildContext context) {
    return ListView.builder(
      itemCount: connectedDevices.length,
      itemBuilder: (BuildContext context, int index) {
        final device = connectedDevices[index];
        final deviceColor = Color.fromRGBO(
          device.info.color.r,
          device.info.color.g,
          device.info.color.b,
          1,
        );
        return Card(
          color: _isSelected(device)
              ? Theme.of(context).colorScheme.tertiaryContainer
              : null,
          margin: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
          child: ListTile(
            leading: CircleAvatar(backgroundColor: deviceColor),
            title: Text(device.info.name),
            subtitle: Text(device.info.type),
            onLongPress: () => _showDeviceDetailsSheet(context, device),
            onTap: () =>
                _isSelected(device) ? _unselect(device) : _select(device),
          ),
        );
      },
    );
  }

  bool _isSelected(dalat.ConnectedDevice dev) {
    for (final device in choosedDevices) {
      if (device.info.id == dev.info.id) return true;
    }
    return false;
  }

  void _select(dalat.ConnectedDevice dev) {
    if (_isSelected(dev)) return;
    setState(() {
      var devs = choosedDevices.map((d) => d).toList();
      devs.add(dev);
      choosedDevices = devs;
    });
    widget.onChange(choosedDevices);
  }

  void _unselect(dalat.ConnectedDevice dev) {
    setState(() {
      var devs = choosedDevices.map((d) => d).toList();
      devs.removeWhere((device) => device.info.id == dev.info.id);
      choosedDevices = devs;
    });
    widget.onChange(choosedDevices);
  }
}
