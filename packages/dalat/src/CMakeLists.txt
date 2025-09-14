# The Flutter tooling requires that developers have CMake 3.10 or later
# installed. You should not increase this version, as doing so will cause
# the plugin to fail to compile for some customers of the plugin.
cmake_minimum_required(VERSION 3.10)

project(dalat_library VERSION 0.0.1 LANGUAGES C)

add_library(alat SHARED IMPORTED)
set_target_properties(alat PROPERTIES IMPORTED_LOCATION ${CMAKE_CURRENT_SOURCE_DIR}/libalat.so)


target_compile_definitions(dalat PUBLIC DART_SHARED_LIB)

target_link_libraries(dalat_plugin PRIVATE alat)


if (ANDROID)
  # Support Android 15 16k page size
  target_link_options(dalat PRIVATE "-Wl,-z,max-page-size=16384")
endif()
