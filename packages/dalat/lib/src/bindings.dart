import 'dart:ffi';
import 'dart:io';

import '../libalat.dart';

// This file is an internal implementation detail of the plugin.
//
// It handles loading the dynamic library and creating the FFI bindings
// so that the high-level API classes can use them.

const String _libName = 'alat';

final DynamicLibrary _dylib = () {
  if (Platform.isMacOS || Platform.isIOS) {
    return DynamicLibrary.open('lib$_libName.dylib');
  } else if (Platform.isAndroid || Platform.isLinux) {
    return DynamicLibrary.open('lib$_libName.so');
  } else if (Platform.isWindows) {
    return DynamicLibrary.open('$_libName.dll');
  } else {
    throw UnsupportedError('Unknown platform: ${Platform.operatingSystem}');
  }
}();

/// The bindings to the native functions in [_dylib].
final AlatBindings bindings = AlatBindings(_dylib);
