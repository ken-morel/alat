import 'package:alat/pages/setup/state.dart';
import 'package:flutter/material.dart';

class ServicesHomePage extends StatelessWidget {
  final SetupState setupState;
  const ServicesHomePage({super.key, required this.setupState});
  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Text("Services", style: Theme.of(context).textTheme.headlineLarge),
        SizedBox(height: 30),
        Text(
          "Services are the different features of your device you want to make available to connected devices. You can disable them at any time.",
          style: Theme.of(context).textTheme.bodyLarge,
        ),
        SizedBox(height: 10),
        Text(
          "The settings are already in sensible defaults which should just fit you, but read attentively so you know what you expose to others.",
          style: Theme.of(context).textTheme.bodyLarge,
        ),
      ],
    );
  }
}
