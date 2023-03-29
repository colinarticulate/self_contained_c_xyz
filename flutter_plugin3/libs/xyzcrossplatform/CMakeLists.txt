if(ANDROID)
    add_library(xyzcrossplatform SHARED 
    android/android.c
    android/fmemopen.c 
    android/getline.c)
    
    set_property(TARGET xyzcrossplatform PROPERTY C_STANDARD 11)

    target_include_directories(
    xyzcrossplatform PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}
    xyzcrossplatform PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/android
)
elseif(CMAKE_SYSTEM_NAME STREQUAL "Darwin")
    add_library(xyzcrossplatform SHARED macos/fmemopen.c)
    set_property(TARGET xyzcrossplatform PROPERTY C_STANDARD 11)

    target_include_directories(
    xyzcrossplatform PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}
    xyzcrossplatform PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/macos
)
elseif(CMAKE_SYSTEM_NAME STREQUAL "iOS")
    add_library(xyzcrossplatform SHARED ios/fmemopen.c)
    set_property(TARGET xyzcrossplatform PROPERTY C_STANDARD 11)

    target_include_directories(
    xyzcrossplatform PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}
    xyzcrossplatform PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/ios
)
else()
    add_library(xyzcrossplatform SHARED linux/fmemopen.c linux/getline.c)
    set_property(TARGET xyzcrossplatform PROPERTY C_STANDARD 11)

    target_include_directories(
    xyzcrossplatform PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}
    xyzcrossplatform PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/linux
)
endif()
# set_property(TARGET xyzcrossplatform PROPERTY C_STANDARD 11)

# target_include_directories(
#     xyzcrossplatform PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}
#     # xyzcrossplatform INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}
#     # xyzcrossplatform PRIVATE ${CMAKE_CURRENT_SOURCE_DIR} #--->fatal error: crossplatform.h: No such file or directory
# )
