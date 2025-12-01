import 'dart:async';

import 'package:alat/components/devicebatteryview.dart';
import 'package:alat/pages/dashboard/base.dart';
import 'package:alat/pages/dashboard/deviceview.dart';
import 'package:alat/state.dart';
import 'package:flutter/material.dart';
import 'package:alat/l10n/app_localizations.dart';
import 'package:dalat/dalat.dart' as dalat;
import 'package:provider/provider.dart';

class DashboardPage extends DashboardBase {
  const DashboardPage({super.key});
  @override
  AppBar buildAppBar(BuildContext context) {
    return AppBar(title: Text(AppLocalizations.of(context)!.dashboard));
  }

  @override
  Widget buildContent(BuildContext context) {
    return Column(
      children: [
        Center(
          child: Text(
            AppLocalizations.of(context)!.activeDevices,
            style: Theme.of(context).textTheme.headlineLarge,
          ),
        ),
        SizedBox(height: 10),
        _ConnectedDevicesList(),
        SizedBox(height: 10),
        FilledButton.tonal(
          onPressed: () =>
              Navigator.of(context).pushReplacementNamed("/dashboard/pair"),

          child: Text(AppLocalizations.of(context)!.connectANewDevice),
          // automatically changed to localization, thanks to arb-util
        ),
      ],
    );
  }
}

class _ConnectedDevicesList extends StatefulWidget {
  @override
  State<_ConnectedDevicesList> createState() => _ConnectedDevicesListState();
}

class _ConnectedDevicesListState extends State<_ConnectedDevicesList> {
  List<dalat.ConnectedDevice> connecteDevices = [];
  late Timer timer;

  @override
  void initState() {
    final AppState appState = context.read();
    timer = Timer.periodic(Duration(seconds: 1), (_) {
      appState.node?.getConnectedDevices().then((devices) {
        setState(() {
          connecteDevices = devices;
        });
      });
    });
    super.initState();
  }

  @override
  void dispose() {
    timer.cancel();
    super.dispose();
  }

  Widget _buildNoConnectedWidget(BuildContext context) {
    return Card(
      child: Padding(
        padding: EdgeInsets.symmetric(horizontal: 30, vertical: 20),
        child: Text(AppLocalizations.of(context)!.noActiveDevice),
      ),
    );
  }

  Widget _buildConnectedDevicesView(BuildContext context) {
    return SingleChildScrollView(
      child: Column(
        children: connecteDevices.map((device) {
          final deviceColor = Color.fromRGBO(
            device.info.color.r,
            device.info.color.g,
            device.info.color.b,
            1,
          );
          return Card(
            margin: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
            child: ListTile(
              leading: CircleAvatar(
                backgroundColor: deviceColor,
                child: DeviceBatteryView(connectedDevice: device),
              ),
              title: Text(device.info.name),
              subtitle: Text(device.info.type),
              onTap: () => Navigator.of(context).push(
                MaterialPageRoute(
                  builder: (context) => DeviceView(device: device),
                ),
              ),
            ),
          );
        }).toList(),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return connecteDevices.isEmpty
        ? _buildNoConnectedWidget(context)
        : _buildConnectedDevicesView(context);
  }
}
