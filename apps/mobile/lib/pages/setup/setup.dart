import 'package:alat/pages/setup/complete.dart';
import 'package:alat/pages/setup/device.dart';
import 'package:alat/pages/setup/home.dart';
import 'package:alat/pages/setup/services/filesend.dart';
import 'package:alat/pages/setup/services/services.dart';
import 'package:alat/pages/setup/services/sysinfo.dart';
import 'package:alat/pages/setup/state.dart';
import 'package:alat/state.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:alat/l10n/app_localizations.dart';

class SetupAssistantPageView extends StatelessWidget {
  final controller = PageController();
  SetupAssistantPageView({super.key});
  Widget _wrap(
    BuildContext context,
    SetupState state,
    Widget content,
    bool prev,
    bool next,
  ) {
    return Padding(
      padding: EdgeInsetsGeometry.only(top: 30, left: 30, right: 30),
      child: SingleChildScrollView(
        child: Column(
          children: [
            content,
            SizedBox(height: 50),
            _buildNav(context, state, prev, next),
          ],
        ),
      ),
    );
  }

  Widget _buildNav(
    BuildContext context,
    SetupState state,
    bool prev,
    bool next,
  ) {
    return Row(
      mainAxisAlignment: prev && next
          ? MainAxisAlignment.spaceBetween
          : prev
          ? MainAxisAlignment.spaceBetween
          : MainAxisAlignment.end,
      children: [
        if (prev)
          FilledButton.tonal(
            onPressed: () {
              state.prev();
            },
            child: Text(AppLocalizations.of(context)!.previous),
          ),
        if (next)
          FilledButton.tonal(
            onPressed: () {
              state.next();
            },
            child: Text(AppLocalizations.of(context)!.next),
          ),
      ],
    );
  }

  Widget buildContent(BuildContext context) {
    final state = SetupState(controller, context.read<AppState>());
    final pages = [
      _wrap(context, state, SetupHome(setupState: state), false, true),
      _wrap(context, state, SetupDevice(setupState: state), true, true),
      _wrap(context, state, ServicesHomePage(setupState: state), true, true),
      _wrap(context, state, SysInfoSetupPage(setupState: state), true, true),
      _wrap(context, state, FileSendSetupPage(setupState: state), true, true),
      _wrap(context, state, SetupCompletePage(setupState: state), true, false),
    ];
    return PageView.builder(
      controller: controller,
      itemBuilder: (BuildContext context, int index) {
        return pages[index % pages.length];
      },
      itemCount: pages.length,
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text(AppLocalizations.of(context)!.alatSetup)),
      body: buildContent(context),
    );
  }
}
