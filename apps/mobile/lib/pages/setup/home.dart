import 'package:alat/pages/setup/state.dart';
import 'package:flutter/material.dart';
import 'package:alat/l10n/app_localizations.dart';

class SetupHome extends StatelessWidget {
  final SetupState setupState;
  const SetupHome({super.key, required this.setupState});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Text(
          AppLocalizations.of(context)!.welcomeToAlat,
          style: Theme.of(context).textTheme.headlineLarge,
        ),
        SizedBox(height: 10),
        Text(
          AppLocalizations.of(
            context,
          )!.weAreGoingToGoThroughOutTheProcessOfSettingUpYourDevice,
        ),
        SizedBox(height: 30),
      ],
    );
  }
}
