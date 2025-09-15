# The Flutter tooling requires that developers have CMake 3.10 or later
# installed. You should not increase this version, as doing so will cause
# the plugin to fail to compile for some customers of the plugin.
cmake_minimum_required(VERSION 3.10)

project(dalat_library VERSION 0.0.1 LANGUAGES C)

if(NOT TARGET alat)
  add_library(alat SHARED IMPORTED)
  set_target_properties(alat PROPERTIES IMPORTED_LOCATION ${CMAKE_CURRENT_SOURCE_DIR}/libalat.so)
endif()

# The linking to the actual plugin is handled in the platform-specific CMakeLists.txt
# (e.g., linux/CMakeLists.txt)
