import 'dart:ffi';
import 'dart:io';

import '../libalat.dart';

final DynamicLibrary _dylib = () {
  if (Platform.isMacOS || Platform.isIOS) {
    return DynamicLibrary.open('libalat.dylib');
  } else if (Platform.isAndroid || Platform.isLinux) {
    return DynamicLibrary.open('libalat.so');
  } else if (Platform.isWindows) {
    return DynamicLibrary.open('alat.dll');
  } else {
    throw UnsupportedError('Unknown platform: ${Platform.operatingSystem}');
  }
}();

/// The bindings to the native functions in [_dylib].
final AlatBindings bindings = AlatBindings(_dylib);
