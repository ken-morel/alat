import 'package:alat/components/sendfiles.dart';
import 'package:alat/pages/dashboard/base.dart';
import 'package:dalat/dalat.dart' as dalat;
import 'package:flutter/material.dart';

class SendFilesPage extends DashboardBase {
  final List<dalat.ConnectedDevice> devices;
  final List<String> files;
  const SendFilesPage({
    super.key,
    this.devices = const [],
    this.files = const [],
  });
  @override
  AppBar buildAppBar(BuildContext context) {
    return AppBar(title: Text("Send files"));
  }

  @override
  Widget buildContent(BuildContext context) {
    return SendFilesWidget(selectedDevices: devices, selectedFiles: files);
  }
}
