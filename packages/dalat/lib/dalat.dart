import 'dart:ffi';
import 'dart:io';

import 'libalat.dart';

const String _libName = 'alat';

/// The dynamic library in which the symbols for [AlatBindings] can be found.
final DynamicLibrary _dylib = () {
  if (Platform.isMacOS || Platform.isIOS) {
    // The FFI plugin template is not used, so we don't use the framework linking.
    // Instead, we load the dylib directly.
    return DynamicLibrary.open('lib$_libName.dylib');
  }
  if (Platform.isAndroid || Platform.isLinux) {
    return DynamicLibrary.open('lib$_libName.so');
  }
  if (Platform.isWindows) {
    return DynamicLibrary.open('$_libName.dll');
  }
  throw UnsupportedError('Unknown platform: ${Platform.operatingSystem}');
}();

/// The bindings to the native functions in [_dylib].
final AlatBindings bindings = AlatBindings(_dylib);
