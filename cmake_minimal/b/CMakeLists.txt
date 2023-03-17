add_library(b SHARED b.c)

target_include_directories(
    b PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}
)
