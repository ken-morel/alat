import 'package:flutter/material.dart';
import 'package:dalat/dalat.dart' as dalat;

class AlatStatusWidget extends StatefulWidget {
  final dalat.AlatInstance node;
  const AlatStatusWidget({super.key, required this.node});
  @override
  State<AlatStatusWidget> createState() => _AlatStatusWidgetState();
}

class _AlatStatusWidgetState extends State<AlatStatusWidget> {
  @override
  Widget build(BuildContext context) {
    return Text("Widget");
  }
}
