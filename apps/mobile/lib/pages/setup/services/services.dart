import 'package:alat/pages/setup/state.dart';
import 'package:flutter/material.dart';
import 'package:alat/l10n/app_localizations.dart';

class ServicesHomePage extends StatelessWidget {
  final SetupState setupState;
  const ServicesHomePage({super.key, required this.setupState});
  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Text(
          AppLocalizations.of(context)!.services,
          style: Theme.of(context).textTheme.headlineLarge,
        ),
        SizedBox(height: 30),
        Text(
          AppLocalizations.of(
            context,
          )!.servicesAreTheDifferentFeaturesOfYourDeviceYouWantToMakeAvailableToConnectedDevicesYouCanDisableThemAtAnyTime,
          style: Theme.of(context).textTheme.bodyLarge,
        ),
        SizedBox(height: 10),
        Text(
          AppLocalizations.of(
            context,
          )!.theSettingsAreAlreadyInSensibleDefaultsWhichShouldJustFitYouButReadAttentivelySoYouKnowWhatYouExposeToOthers,
          style: Theme.of(context).textTheme.bodyLarge,
        ),
      ],
    );
  }
}
