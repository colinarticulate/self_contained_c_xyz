# The Flutter tooling requires that developers have CMake 3.10 or later
# installed. You should not increase this version, as doing so will cause
# the plugin to fail to compile for some customers of the plugin.
cmake_minimum_required(VERSION 3.10)

project(ffigen_cmake_library VERSION 0.0.1 LANGUAGES C)

add_library(ffigen_cmake SHARED
  "ffigen_cmake.c"
)

set_target_properties(ffigen_cmake PROPERTIES
  PUBLIC_HEADER ffigen_cmake.h
  OUTPUT_NAME "ffigen_cmake"
)

target_compile_definitions(ffigen_cmake PUBLIC DART_SHARED_LIB)
