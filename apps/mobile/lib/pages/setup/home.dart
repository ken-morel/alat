import "package:alat/pages/setup/state.dart";
import "package:flutter/material.dart";

class SetupHome extends StatelessWidget {
  final SetupState setupState;
  const SetupHome({super.key, required this.setupState});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Text(
          "Welcome to alat",
          style: Theme.of(context).textTheme.headlineLarge,
        ),
        SizedBox(height: 10),
        Text(
          "We are going to go through out the process of setting up your device",
        ),
        SizedBox(height: 30),
        Row(
          mainAxisAlignment: MainAxisAlignment.end,
          children: [
            FilledButton.tonal(
              onPressed: () {
                setupState.next();
              },
              child: Text("Next"),
            ),
          ],
        ),
      ],
    );
  }
}
