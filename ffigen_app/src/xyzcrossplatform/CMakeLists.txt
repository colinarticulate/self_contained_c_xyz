add_library(xyzcrossplatform SHARED crossplatform.c)
set_property(TARGET xyzcrossplatform PROPERTY C_STANDARD 11)


target_include_directories(
    xyzcrossplatform PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}
    # xyzcrossplatform INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}
    # xyzcrossplatform PRIVATE ${CMAKE_CURRENT_SOURCE_DIR} #--->fatal error: crossplatform.h: No such file or directory
)
