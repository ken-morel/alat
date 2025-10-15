import 'package:alat/components/founddevices.dart';
import 'package:alat/pages/dashboard/base.dart';
import 'package:flutter/material.dart';

class PairDevicePage extends DashboardBase {
  const PairDevicePage({super.key});
  @override
  AppBar buildAppBar(BuildContext context) {
    return AppBar(title: Text("Connect a device"));
  }

  @override
  Widget buildContent(BuildContext context) {
    return Column(
      children: [
        Center(
          child: Text(
            "Found devices",
            style: Theme.of(context).textTheme.headlineLarge,
          ),
        ),
        SizedBox(
          height: 300,
          child: FoundDevicesList(onConnectionUserRequest: (device) {}),
        ),
      ],
    );
  }
}
