cmake_minimum_required(VERSION 3.15)

project(ps_plus)#PROJECT_SOURCE_DIR#CMAKE_SOURCE_DIR

# target_compile_features(pocketsphinx_compiler_flags INTERFACE c_std_11)
# target_compile_features(sphinxbase_compiler_flags INTERFACE c_std_11)

set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR})

add_subdirectory(xyzpocketsphinx)
add_subdirectory(xyzsphinxbase)
# list(APPEND EXTRA_LIBS xyzpocketsphinx xyzsphinxbase xyzsphinxad)

target_link_libraries(xyzsphinxad PUBLIC xyzsphinxbase)
target_link_libraries(xyzpocketsphinx PUBLIC xyzsphinxbase xyzsphinxad)
target_include_directories(xyzpocketsphinx PUBLIC
                           "${PROJECT_BINARY_DIR}"
                           )

add_executable(ps_plus ps_plus.cpp )

target_link_libraries(ps_plus PUBLIC xyzpocketsphinx)

# target_link_libraries(Tutorial PUBLIC ${EXTRA_LIBS})

target_include_directories(
ps_plus PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}
ps_plus PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/aside_ps_library
ps_plus PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/_xyzsphinxbase
ps_plus PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/_xyzsphinxbase/fe
)
