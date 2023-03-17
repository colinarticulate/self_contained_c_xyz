cmake_minimum_required(VERSION 3.20)

project(ps_plus)#PROJECT_SOURCE_DIR#CMAKE_SOURCE_DIR

# target_compile_features(pocketsphinx_compiler_flags INTERFACE c_std_11)
# target_compile_features(sphinxbase_compiler_flags INTERFACE c_std_11)

#set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${PROJECT_SOURCE_DIR}/bin) #Final output of the executable
# set(CMAKE_CXX_STANDARD_REQUIRED ON)
# set(CMAKE_CXX_EXTENSIONS OFF) #Set this to ON if you want to use GNU++
# set(CMAKE_CXX_STANDARD 11)

# set(CMAKE_C_STANDARD_REQUIRED ON)
# set(CMAKE_C_EXTENSIONS OFF) #Set this to ON if you want to use GNU++
# set(CMAKE_C_STANDARD 11)

add_subdirectory(xyzpocketsphinx)
add_subdirectory(xyzsphinxbase)
add_subdirectory(xyzcrossplatform)


add_definitions(-DHAVE_CONFIG_H)
add_compile_options(-O3)

target_link_libraries(xyzsphinxbase PUBLIC xyzcrossplatform)
target_link_libraries(xyzsphinxad PUBLIC xyzsphinxbase)
target_link_libraries(xyzpocketsphinx PUBLIC xyzsphinxbase xyzsphinxad)

target_include_directories(xyzpocketsphinx PUBLIC "${PROJECT_BINARY_DIR}")
# target_include_directories(xyzcrossplatform PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/crossplatform)

#add_executable(ps_plus ps_plus.cpp )
#add_executable(ps_plus c_demo.cpp )
add_library(ps_plus
  SHARED
  ./c_demo.cpp
  exports.def
)

set_target_properties(ps_plus PROPERTIES
  PUBLIC_HEADER c_demo_ffi.h
  OUTPUT_NAME "ffigen_app"
)

target_compile_definitions(ps_plus PUBLIC DART_SHARED_LIB)

#set_property(TARGET ps_plus PROPERTY CXX_STANDARD 11)
# target_compile_features(ps_plus PUBLIC cxx_std_11)

#target_link_libraries(ps_plus PUBLIC xyzpocketsphinx)
target_link_libraries(ps_plus PUBLIC  xyzpocketsphinx xyzcrossplatform)


target_include_directories(
ps_plus PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}
ps_plus PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/aside_ps_library
ps_plus PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/_xyzsphinxbase
ps_plus PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/_xyzsphinxbase/fe
# ps_plus PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/xyzcrossplatform
)
