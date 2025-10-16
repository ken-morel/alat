import 'package:alat/components/deviceicon.dart';
import 'package:alat/pages/dashboard/base.dart';
import 'package:dalat/dalat.dart' as dalat;
import 'package:flutter/material.dart';

class DeviceView extends DashboardBase {
  final dalat.ConnectedDevice device;
  const DeviceView({super.key, required this.device});
  @override
  AppBar buildAppBar(BuildContext context) {
    return AppBar(leading: BackButton());
  }

  @override
  Widget buildContent(BuildContext context) {
    // device.info.name
    // device.info.color
    // device.info.id
    // device.info.type
    // device.pairedDevice.certificate
    // device.pairedDevice.token
    // device.port
    return Column(
      children: [
        DeviceIcon(deviceType: device.info.type, color: device.info.color),
        Text(
          device.info.name,
          style: Theme.of(context).textTheme.headlineMedium,
        ),
      ],
    );
  }
}
