import 'package:dalat/dalat.dart';
import 'package:dalat/src/bindings.dart';
import 'package:dalat/src/helpers.dart';

mixin InstanceWebShare on InstanceHelpers {
  Future<WebshareStatus> getWebshareStatus() {
    return jsonGetterHelper(
      bindings.wehsbare_get_status_json,
      WebshareStatus.fromJson,
    );
  }

  Future<void> addSharedFiles(List<String> files) {
    return setterHelper(files.join(";"), bindings.webshare_add_shared_files);
  }

  Future<void> removeSharedFilesByPaths(List<String> files) {
    return setterHelper(
      files.join(";"),
      bindings.webshare_remove_shared_files_by_paths,
    );
  }

  Future<void> removeSharedFilesByUUIDs(List<String> files) {
    return setterHelper(
      files.join(";"),
      bindings.webshare_remove_shared_files_by_uuids,
    );
  }

  Future<void> clearSharedFiles() {
    return helper(bindings.webshare_clear_shared_files);
  }

  Future<List<SharedFile>> getSharedFiles() {
    return jsonListGetterHelper(
      bindings.webshare_get_shared_files_json,
      SharedFile.fromJson,
    );
  }

  Future<String> getWebsharePasscode() {
    return getterHelper(bindings.webshare_get_passcode);
  }

  Future<void> setWebsharePasscode(String passcode) {
    return setterHelper(passcode, bindings.webshare_set_passcode);
  }

  Future<void> stopWebshare() {
    return helper(bindings.webshare_stop);
  }

  Future<void> startWebshare() {
    return helper(bindings.webshare_start);
  }
}
