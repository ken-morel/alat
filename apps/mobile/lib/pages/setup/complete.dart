import 'package:alat/pages/setup/state.dart';
import 'package:flutter/material.dart';
import 'package:alat/l10n/app_localizations.dart';

class SetupCompletePage extends StatelessWidget {
  final SetupState setupState;
  const SetupCompletePage({super.key, required this.setupState});
  @override
  Widget build(BuildContext context) {
    setupState.appState.completeSetup();
    return Column(
      children: [
        Icon(Icons.check_circle_rounded, color: Colors.greenAccent, size: 200),
        Text(
          AppLocalizations.of(context)!.setupComplete,
          style: Theme.of(context).textTheme.headlineLarge,
        ),
        SizedBox(height: 30),
        Text(
          AppLocalizations.of(
            context,
          )!.yourDeviceIsCompletelySetupAndNowReadyToStart,
          style: Theme.of(context).textTheme.bodyLarge,
        ),
        SizedBox(height: 20),
        FilledButton(
          onPressed: () {
            Navigator.of(
              context,
            ).pushNamedAndRemoveUntil("/dashboard", (_) => false);
          },
          child: Text("Open dashboard"),
        ),
      ],
    );
  }
}
