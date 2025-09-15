import 'package:alat/pages/setup/device.dart';
import 'package:alat/pages/setup/home.dart';
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
    bool nav,
  ) {
    return Padding(
      padding: EdgeInsetsGeometry.only(top: 30, left: 30, right: 30),
      child: SingleChildScrollView(
        child: Column(
          children: [
            content,
            SizedBox(height: 50),
            if (nav) _buildNav(context, state),
          ],
        ),
      ),
    );
  }

  Widget _buildNav(BuildContext context, SetupState state) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        FilledButton.tonal(
          onPressed: () {
            state.prev();
          },
          child: Text(AppLocalizations.of(context)!.previous),
        ),
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
    return PageView.builder(
      controller: controller,
      itemBuilder: (BuildContext context, int index) {
        switch (index) {
          case 0:
            return _wrap(context, state, SetupHome(setupState: state), false);
          case 1:
            return _wrap(context, state, SetupDevice(setupState: state), true);
          case 2:
            return _wrap(
              context,
              state,
              ServicesHomePage(setupState: state),
              true,
            );
          case 3:
            return _wrap(
              context,
              state,
              SysInfoSetupPage(setupState: state),
              true,
            );
          default:
            return _wrap(context, state, SetupHome(setupState: state), true);
        }
      },
      itemCount: 4,
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
