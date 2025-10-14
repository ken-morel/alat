import 'dart:async';

import 'package:alat/state.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter/widgets.dart';
import 'package:dalat/dalat.dart' as dalat;
import 'package:provider/provider.dart';

class FoundDevicesList extends StatefulWidget {
  @override
  State<FoundDevicesList> createState() => _FoundDeviceListState();
}

class _FoundDeviceListState extends State<FoundDevicesList> {
  List<dalat.FoundDevice> foundDevices = [];
  late Timer timer;
  @override
  void initState() {
    final appState = context.read<AppState>();
    timer = Timer.periodic(
      const Duration(seconds: 1),
      (_) => () async {
        print("Updating list");
        final devices = await appState.node?.getFoundDevices();
        if (devices != null) {
          setState(() {
            foundDevices = devices;
          });
        }
      },
    );
    super.initState();
  }

  @override
  void dispose() {
    timer.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return foundDevices.isEmpty
        ? Text("No found devices, for now")
        : ListView.builder(
            itemCount: foundDevices.length,
            itemBuilder: (BuildContext context, int index) {
              final device = foundDevices[index];
              device.info.name
              device.info.id // 40 characters long id
              device.info.color.(name, hex, r, g, b)
              device.info.type( dalat.DeviceType.(unspecified, mobile, desktop))
              device.ip (Uint8List)
              device.port;
              return ListTile(title: Text(device.info.name));
            },
          );
  }
}
