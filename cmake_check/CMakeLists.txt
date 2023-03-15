cmake_minimum_required(VERSION 3.20)

project(hello)

set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${PROJECT_SOURCE_DIR}/bin)

add_executable(hello source.c )
set_property(TARGET hello PROPERTY C_STANDARD 11)