import 'package:alat/components/alatstatus.dart';
import 'package:alat/l10n/app_localizations.dart';
import 'package:flutter/material.dart';
import 'package:alat/state.dart';
import 'package:provider/provider.dart';

class DashboardBase extends StatelessWidget {
  const DashboardBase({super.key});

  Widget buildContent(BuildContext context) {
    return const Text("Nothing here");
  }

  Widget _buildLink(BuildContext context, String label, String name) {
    return TextButton(
      onPressed: () => Navigator.of(context).pushReplacementNamed(name),
      child: Text(label),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      drawer: Drawer(
        child: ListView(
          children: [
            AlatStatusWidget(node: context.read<AppState>().node!),
            Divider(),
            _buildLink(context, "Dashboard", "/dashboard"),
            Divider(),
            _buildLink(context, "Connect a device", "/dashboard/pair"),
          ],
        ),
      ),
      appBar: AppBar(
        title: Text(AppLocalizations.of(context)!.dashboard),
        leading: Builder(
          builder: (BuildContext context) => IconButton(
            onPressed: () {
              Scaffold.of(context).openDrawer();
            },
            icon: Icon(Icons.menu),
          ),
        ),
      ),
      body: buildContent(context),
    );
  }
}
