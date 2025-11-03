import 'package:flutter/material.dart';
import 'package:file_picker/file_picker.dart';

class FilesSelectionComponent extends StatefulWidget {
  final Future<void> Function(List<PlatformFile>) onChange;
  const FilesSelectionComponent({super.key, required this.onChange});
  @override
  State<StatefulWidget> createState() => _FilesSelectionComponentState();
}

String formatSize(int bytes) {
  if (bytes < 1024) {
    return "$bytes Bytes";
  } else if (bytes < 1024 * 1024) {
    return "${(bytes / 1024).toStringAsPrecision(3)} Kilo Bytes";
  } else if (bytes < 1024 * 1024 * 1024) {
    return "${(bytes / (1024 * 1024)).toStringAsPrecision(3)} Mega Bytes";
  } else {
    return "${(bytes / (1024 * 1024 * 1024)).toStringAsPrecision(3)} Giga Bytes";
  }
}

class _FilesSelectionComponentState extends State<FilesSelectionComponent> {
  List<PlatformFile> selectedFiles = [];
  String? error;
  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Expanded(
          child: Card(
            margin: const EdgeInsets.symmetric(horizontal: 30, vertical: 10),
            child: Column(
              children: [
                selectedFiles.isNotEmpty
                    ? Expanded(
                        child: ListView.builder(
                          itemCount: selectedFiles.length,
                          itemBuilder: (context, index) {
                            final file = selectedFiles[index];
                            return ListTile(
                              leading: const Icon(Icons.file_upload_rounded),
                              title: Text(file.name),
                              subtitle: Text(formatSize(file.size)),
                              trailing: IconButton(
                                onPressed: () {
                                  setState(() {
                                    selectedFiles.remove(file);
                                  });
                                  widget.onChange(selectedFiles);
                                },
                                icon: Icon(Icons.cancel_rounded),
                              ),
                            );
                          },
                        ),
                      )
                    : const Expanded(
                        child: Center(child: Text("No file selected")),
                      ),
                Padding(
                  padding: EdgeInsets.only(left: 10, right: 10, bottom: 10),
                  child: SizedBox(
                    width: double.infinity,
                    child: FilledButton.tonal(
                      onPressed: () {
                        FilePicker.platform.pickFiles(allowMultiple: true).then(
                          (files) {
                            if (files == null) return;
                            setState(() {
                              selectedFiles.addAll(files.files);
                            });
                            widget.onChange(selectedFiles);
                          },
                        );
                      },
                      child: const Text("Add files"),
                    ),
                  ),
                ),
              ],
            ),
          ),
        ),
        if (error != null)
          Text(
            error!,
            style: Theme.of(context).textTheme.labelMedium?.copyWith(
              color: Theme.of(context).colorScheme.error,
            ),
          ),

        SizedBox(height: 5),
      ],
    );
  }
}
