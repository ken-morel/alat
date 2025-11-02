import 'package:alat/components/deviceselection.dart';
import 'package:alat/components/filesselection.dart';
import 'package:alat/state.dart';
import 'package:flutter/material.dart';
import 'package:dalat/dalat.dart' as dalat;
import 'package:provider/provider.dart';

class SendFilesWidget extends StatefulWidget {
  final List<dalat.ConnectedDevice> selectedDevices;
  final List<String> selectedFiles;
  const SendFilesWidget({
    super.key,
    this.selectedDevices = const [],
    this.selectedFiles = const [],
  });
  @override
  State<StatefulWidget> createState() => _SendFilesWidgetState();
}

class _SendFilesWidgetState extends State<SendFilesWidget> {
  late List<dalat.ConnectedDevice> selectedDevices;
  late List<String> selectedFiles;
  @override
  void initState() {
    selectedDevices = widget.selectedDevices;
    selectedFiles = widget.selectedFiles;
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    final appState = context.read<AppState>();
    return Column(
      children: [
        Text("Select Files", style: Theme.of(context).textTheme.displaySmall),
        Expanded(
          child: FilesSelectionComponent(
            onChange: (files) async {
              setState(() {
                selectedFiles = files.map((f) => f.path).nonNulls.toList();
              });
            },
          ),
        ),
        Text("Select Devices", style: Theme.of(context).textTheme.displaySmall),
        Expanded(
          child: Padding(
            padding: EdgeInsets.symmetric(horizontal: 20, vertical: 10),
            child: ConnecedDevicesSelection(
              selecedDevices: selectedDevices,
              onChange: (devices) {
                setState(() {
                  selectedDevices = devices;
                });
              },
            ),
          ),
        ),
        SizedBox(
          width: double.infinity,
          child: Padding(
            padding: EdgeInsets.symmetric(horizontal: 20, vertical: 20),
            child: FilledButton(
              onPressed: selectedFiles.length * selectedDevices.length == 0
                  ? null
                  : () {
                      for (final device in selectedDevices) {
                        appState.node?.querySendFilesToDevice(
                          device.info.id,
                          selectedFiles,
                        );
                        Navigator.of(
                          context,
                        ).pushReplacementNamed("/dashboard");
                      }
                    },
              child: Text(
                "Send ${selectedFiles.length} file(s) to ${selectedDevices.length} device(s)",
              ),
            ),
          ),
        ),
      ],
    );
  }
}

// Divider(),
//         Padding(
//           padding: EdgeInsetsGeometry.symmetric(horizontal: 15, vertical: 10),
//           child: SizedBox(
//             width: double.infinity,
//             child: FilledButton(
//               onPressed: () {
//                 () async {
//                   try {
//                     widget.onSubmit(selectedFiles);
//                     setState(() {
//                       error = null;
//                       selectedFiles.removeWhere((_) => true);
//                       Navigator.of(context).pop();
//                     });
//                   } catch (e) {
//                     setState(() {
//                       error = e.toString();
//                     });
//                   }
//                 }();
//               },
//               child: Text("Send files"),
//             ),
//           ),
//         ),
