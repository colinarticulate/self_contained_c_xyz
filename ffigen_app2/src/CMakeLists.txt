# The Flutter tooling requires that developers have CMake 3.10 or later
# installed. You should not increase this version, as doing so will cause
# the plugin to fail to compile for some customers of the plugin.
cmake_minimum_required(VERSION 3.10)

project(ffigen_app2_library VERSION 0.0.1 LANGUAGES C)

add_library(ffigen_app2 SHARED
  "ffigen_app2.c"
)

set_target_properties(ffigen_app2 PROPERTIES
  PUBLIC_HEADER ffigen_app2.h
  OUTPUT_NAME "ffigen_app2"
)

target_compile_definitions(ffigen_app2 PUBLIC DART_SHARED_LIB)
