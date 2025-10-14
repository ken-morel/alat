import 'package:alat/pages/dashboard/base.dart';
import 'package:flutter/material.dart';
import 'package:alat/l10n/app_localizations.dart';

class DashboardPage extends DashboardBase {
  const DashboardPage({super.key});

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
        SizedBox(height: 50),
        FilledButton.tonal(
          onPressed: () {},
          child: Text("Connect a new device"),
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
  @override
  Widget build(BuildContext context) {
    return Text("Nothing for now");
  }
}
