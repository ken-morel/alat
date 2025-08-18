import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:alat/src/l10n/app_localizations.dart';
import 'package:alat/src/providers/device_provider.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;
    final deviceProvider = context.watch<DeviceProvider>();

    return Scaffold(
      appBar: AppBar(
        title: Text(l10n.homeTitle),
      ),
      body: Center(
        child: Column(
          children: [
            Padding(
              padding: const EdgeInsets.all(16.0),
              child: ElevatedButton(
                onPressed: deviceProvider.isLoading
                    ? null
                    : () => context.read<DeviceProvider>().searchDevices(),
                child: deviceProvider.isLoading
                    ? const CircularProgressIndicator()
                    : const Text('Scan for Devices'),
              ),
            ),
            Expanded(
              child: ListView.builder(
                itemCount: deviceProvider.devices.length,
                itemBuilder: (context, index) {
                  final device = deviceProvider.devices[index];
                  return ListTile(
                    title: Text(device['name'] ?? 'Unknown Device'),
                    subtitle: Text(device['address']?['ip'] ?? 'Unknown IP'),
                  );
                },
              ),
            ),
          ],
        ),
      ),
    );
  }
}