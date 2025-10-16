import 'package:alat/components/devicebatteryview.dart';
import 'package:alat/components/deviceicon.dart';
import 'package:alat/pages/dashboard/base.dart';
import 'package:alat/pages/dashboard/devicefilesend.dart';
import 'package:dalat/dalat.dart' as dalat;
import 'package:flutter/material.dart';

class DeviceView extends DashboardBase {
  final dalat.ConnectedDevice device;
  const DeviceView({super.key, required this.device});

  @override
  AppBar buildAppBar(BuildContext context) {
    return AppBar(
      leading: const BackButton(),
      title: Text(device.info.name),
      elevation: 0,
      backgroundColor: Colors.transparent,
    );
  }

  @override
  Widget buildContent(BuildContext context) {
    final textTheme = Theme.of(context).textTheme;

    return ListView(
      padding: const EdgeInsets.all(16.0),
      children: [
        // --- Header Section ---
        Column(
          children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                DeviceIcon(
                  deviceType: device.info.type,
                  color: device.info.color,
                  size: 80,
                ),
                DeviceBatteryView(connectedDevice: device, size: 70),
              ],
            ),
            const SizedBox(height: 16),
            Text(
              device.info.name,
              style: textTheme.headlineLarge,
              textAlign: TextAlign.center,
            ),
          ],
        ),
        SizedBox(height: 10),

        /// here
        const SizedBox(height: 24),
        // --- Actions Section ---
        Card(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              ListTile(
                leading: const Icon(Icons.upload_file_outlined),
                title: const Text('Send File'),
                subtitle: const Text('Transfer files to this device'),
                onTap: () {
                  Navigator.of(context).push(
                    MaterialPageRoute(
                      builder: (_) =>
                          DeviceFileSendPage(connectedDevice: device),
                    ),
                  );
                },
              ),
              // Add other actions here as ListTiles
            ],
          ),
        ),
        const SizedBox(height: 16),

        // --- Device Info Section ---
        Card(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Padding(
                padding: const EdgeInsets.fromLTRB(16, 16, 16, 8),
                child: Text('Device Details', style: textTheme.titleLarge),
              ),
              Divider(),
              _buildInfoTile(
                context,
                icon: Icons.perm_identity,
                title: 'Device ID',
                subtitle: device.info.id,
              ),
              _buildInfoTile(
                context,
                icon: Icons.lan_outlined,
                title: 'Address',
                subtitle: '${device.ip}:${device.port}',
              ),
              _buildInfoTile(
                context,
                icon: Icons.vpn_key_outlined,
                title: 'Pair Token',
                subtitle: device.pairedDevice.token.toString(),
                isMonospace: true,
              ),
            ],
          ),
        ),
      ],
    );
  }

  Widget _buildInfoTile(
    BuildContext context, {
    required IconData icon,
    required String title,
    required String subtitle,
    bool isMonospace = false,
  }) {
    final textTheme = Theme.of(context).textTheme;
    return ListTile(
      leading: Icon(icon),
      title: Text(title, style: textTheme.titleMedium),
      subtitle: Text(
        subtitle,
        style: isMonospace
            ? textTheme.bodySmall?.copyWith(fontFamily: 'monospace')
            : textTheme.bodyMedium,
      ),
    );
  }
}
