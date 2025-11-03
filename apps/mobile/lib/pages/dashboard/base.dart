import 'package:alat/components/alatstatus.dart';
import 'package:alat/components/themeswitcher.dart';
import 'package:alat/l10n/app_localizations.dart';
import 'package:flutter/material.dart';
import 'package:alat/state.dart';
import 'package:provider/provider.dart';

class DashboardBase extends StatelessWidget {
  const DashboardBase({super.key});

  Widget buildContent(BuildContext context) {
    return const Text("Nothing here");
  }

  AppBar buildAppBar(BuildContext context) {
    return AppBar();
  }

  Widget _buildDrawerItem(
    BuildContext context, {
    required IconData icon,
    required String label,
    required String routeName,
  }) {
    final currentRoute = ModalRoute.of(context)?.settings.name;
    final isSelected = currentRoute == routeName;

    return ListTile(
      leading: Icon(icon),
      title: Text(label),
      selected: isSelected,
      selectedTileColor: Theme.of(context).colorScheme.primary.withAlpha(25),
      onTap: () {
        // Close the drawer
        Navigator.of(context).pop();
        // Navigate if not already on the page
        if (!isSelected) {
          Navigator.of(context).pushReplacementNamed(routeName);
        }
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      drawer: Drawer(
        child: ListView(
          padding: EdgeInsets.zero,
          children: [
            DrawerHeader(
              decoration: BoxDecoration(
                color: Theme.of(context).colorScheme.primaryContainer,
              ),
              margin: EdgeInsets.zero,
              child: AlatStatusWidget(node: context.read<AppState>().node!),
            ),
            _buildDrawerItem(
              context,
              icon: Icons.dashboard_outlined,
              label: AppLocalizations.of(context)!.dashboard,
              routeName: "/dashboard",
            ),
            _buildDrawerItem(
              context,
              icon: Icons.add_circle_outline,
              label: "Connect a device",
              routeName: "/dashboard/pair",
            ),
            const Divider(),
            _buildDrawerItem(
              context,
              icon: Icons.file_copy_rounded,
              label: "Send files",
              routeName: "/sendfiles",
            ),
            _buildDrawerItem(
              context,
              icon: Icons.open_in_browser_rounded,
              label: "Web share",
              routeName: "/webshare",
            ),
            const Divider(),
            const ThemeSwitcher(),
          ],
        ),
      ),
      appBar: buildAppBar(context),

      body: buildContent(context),
    );
  }
}
