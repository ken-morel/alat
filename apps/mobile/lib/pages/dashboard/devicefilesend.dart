import 'package:alat/components/filesselection.dart';
import 'package:alat/pages/dashboard/base.dart';
import 'package:dalat/dalat.dart' as dalat;
import 'package:flutter/material.dart';

class DeviceFileSendPage extends DashboardBase {
  final dalat.ConnectedDevice connectedDevice;
  const DeviceFileSendPage({super.key, required this.connectedDevice});
  @override
  AppBar buildAppBar(BuildContext context) {
    return AppBar(leading: BackButton());
  }

  @override
  Widget buildContent(BuildContext context) {
    return FilesSelectionComponent(onSubmit: (files) {});
  }
}
