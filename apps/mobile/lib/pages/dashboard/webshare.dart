import 'package:alat/components/webshare.dart';
import 'package:alat/pages/dashboard/base.dart';
import 'package:flutter/material.dart';

class WebSharePage extends DashboardBase {
  const WebSharePage({super.key});
  @override
  AppBar buildAppBar(BuildContext context) {
    return AppBar(title: Text("Web share"));
  }

  @override
  Widget buildContent(BuildContext context) => WebShareWidget();
}
