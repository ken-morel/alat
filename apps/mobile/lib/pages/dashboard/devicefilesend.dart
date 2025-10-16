import 'package:alat/components/filesselection.dart';
import 'package:alat/pages/dashboard/base.dart';
import 'package:alat/state.dart';
import 'package:dalat/dalat.dart' as dalat;
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class DeviceFileSendPage extends DashboardBase {
  final dalat.ConnectedDevice connectedDevice;
  const DeviceFileSendPage({super.key, required this.connectedDevice});
  @override
  AppBar buildAppBar(BuildContext context) {
    return AppBar(leading: BackButton(), title: Text("Select files"));
  }

  @override
  Widget buildContent(BuildContext context) {
    final appState = context.read<AppState>();
    return FilesSelectionComponent(
      onSubmit: (files) async {
        final List<String> filePaths = [];
        for (final file in files) {
          if (file.path != null) filePaths.add(file.path!);
        }
        await appState.node!.querySendFilesToDevice(
          connectedDevice.info.id,
          filePaths,
        );
      },
    );
  }
}
