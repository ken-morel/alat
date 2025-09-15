import 'package:alat/components/alatstatus.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:alat/state.dart';

class DashboardPage extends StatelessWidget {
  const DashboardPage({super.key});
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      drawer: Drawer(
        child: ListView(
          children: [
            DrawerHeader(
              decoration: BoxDecoration(color: Colors.blue),
              child: AlatStatusWidget(node: context.read<AppState>().node!),
            ),
          ],
        ),
      ),
      appBar: AppBar(
        title: Text("Dashboard"),
        leading: Builder(
          builder: (BuildContext context) => IconButton(
            onPressed: () {
              Scaffold.of(context).openDrawer();
            },
            icon: Icon(Icons.menu),
          ),
        ),
      ),
      body: Column(
        children: [
          SizedBox(height: 50),
          Center(
            child: Text(
              "Active devices",
              style: Theme.of(context).textTheme.headlineLarge,
            ),
          ),
          SizedBox(height: 10),
          _ConnectedDevicesList(),
        ],
      ),
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
    return Text("NOthing for now");
  }
}
