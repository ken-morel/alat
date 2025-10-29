import 'package:alat/components/filesselection.dart';
import 'package:alat/state.dart';
import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';
import 'package:dalat/dalat.dart' as dalat;
import 'package:provider/provider.dart';

class SendFilesWidget extends StatelessWidget {
  List<dalat.ConnectedDevice> selectedDevices;
  List<String> selectedFiles;
  SendFilesWidget({
    super.key,
    this.selectedDevices = const [],
    this.selectedFiles = const [],
  });
  @override
  Widget build(BuildContext context) {
    final appState = context.read<AppState>();
    return FilesSelectionComponent(
      onChange: (files) async {
        selectedFiles = [];
        for (final file in files) {
          if (file.path != null) selectedFiles.add(file.path!);
        }
      },
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
