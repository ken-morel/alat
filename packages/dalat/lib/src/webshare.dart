import 'package:dalat/dalat.dart';
import 'package:dalat/src/bindings.dart'; // Still needed for direct FFI calls in isolate_helper
import 'package:dalat/src/helpers.dart';
import 'package:dalat/src/helpers.dart';
import 'package:dalat/src/isolate_helper.dart'; // New import

mixin InstanceWebShare on InstanceHelpers {
  Future<WebshareStatus> getWebshareStatus() {
    return jsonGetterHelper(
      'wehsbare_get_status_json',
      WebshareStatus.fromJson,
    );
  }

  Future<void> addSharedFiles(List<String> files) {
    return setterHelper('webshare_add_shared_files', files.join(";"));
  }

  Future<void> removeSharedFilesByPaths(List<String> files) {
    return setterHelper(
      'webshare_remove_shared_files_by_paths',
      files.join(";"),
    );
  }

  Future<void> removeSharedFilesByUUIDs(List<String> files) {
    return setterHelper(
      'webshare_remove_shared_files_by_uuids',
      files.join(";"),
    );
  }

  Future<void> clearSharedFiles() {
    return helper('webshare_clear_shared_files');
  }

  Future<List<SharedFile>> getSharedFiles() {
    return jsonListGetterHelper(
      'webshare_get_shared_files_json',
      SharedFile.fromJson,
    );
  }

  Future<String> getWebsharePasscode() async {
    final result = await FfiIsolate.run('webshare_get_passcode', [handle]) as String?;
    if (result == null) {
      throw Exception('Failed to get passcode from Go core: function returned null.');
    }
    return result;
  }

  Future<void> setWebsharePasscode(String passcode) {
    return setterHelper('webshare_set_passcode', passcode);
  }

  Future<void> stopWebshare() {
    return helper('webshare_stop');
  }

  Future<void> startWebshare() {
    return helper('webshare_start');
  }
}
