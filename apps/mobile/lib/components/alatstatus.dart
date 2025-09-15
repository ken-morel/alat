import 'package:flutter/material.dart';
import 'package:dalat/dalat.dart' as dalat;

class AlatStatusWidget extends StatelessWidget {
  final dalat.AlatInstance node;
  const AlatStatusWidget({super.key, required this.node});

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: node.getNodeStatus(),
      builder: (context, snapshot) {
        if (snapshot.connectionState == ConnectionState.waiting) {
          return CircularProgressIndicator();
        } else if (snapshot.hasError) {
          return Icon(Icons.error);
        } else {
          return Text("Hello: ${snapshot.data}");
        }
      },
    );
  }
}
