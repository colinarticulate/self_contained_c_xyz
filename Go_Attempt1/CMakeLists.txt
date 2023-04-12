cmake_minimum_required(VERSION 3.15)
# project(flutter_pron C CXX Go)
project(flutter_pron)

# Set the C++ library, CGO wrapper, and Go project names
# set(CPP_LIB_NAME mycpplib)
# set(CGO_WRAPPER_NAME mycgowrapper)
# set(GO_PROJ_A_NAME go_proj_A)
# set(GO_PROJ_B_NAME go_proj_B)

# Set the source directories for the C++ library, CGO wrapper, and Go projects
# set(CPP_LIB_SRC_DIR ${CMAKE_SOURCE_DIR}/cpp_lib)
# set(CGO_WRAPPER_SRC_DIR ${CMAKE_SOURCE_DIR}/cgo_wrapper)
# set(GO_PROJ_A_SRC_DIR ${CMAKE_SOURCE_DIR}/go_proj_A)
# set(GO_PROJ_B_SRC_DIR ${CMAKE_SOURCE_DIR}/go_proj_B)

# Compile the C++ library
add_subdirectory(ps_plus)
# add_library(ps_plus SHARED ${CMAKE_SOURCE_DIR}/ps_plus/cpp_lib.cpp)
# add_subdirectory(xyz_plus)
# Compile the CGO wrapper
# add_library(xyz_plus SHARED ${CMAKE_SOURCE_DIR}/xyz_plus)
# target_link_libraries(xyz_plus ps_plus)
# target_include_directories(xyz_plus PRIVATE ${CMAKE_SOURCE_DIR}/xyz_plus)
# set_target_properties(xyz_plus PROPERTIES PREFIX "")

add_custom_command(
    OUTPUT ${CMAKE_BINARY_DIR}/xyz_plus
    COMMAND CGO_LDFLAGS="-L${INSTALL_BUNDLE_LIB_DIR}" LD_LIBRARY_PATH="${INSTALL_BUNDLE_LIB_DIR}" CGO_ENABLED=1 GOOS="${CMAKE_SYSTEM_NAME}" GOARCH="${CMAKE_SYSTEM_PROCESSOR}" GOARM="" go build -buildmode=c-shared -o ${CMAKE_BINARY_DIR}/${PROJECT_NAME} ${SOURCE_DIR}
    DEPENDS ${SOURCE_DIR}/*.go ${DEPENDENCIES}
    WORKING_DIRECTORY ${CMAKE_SOURCE_DIR}
    COMMENT "Building Go project ${PROJECT_NAME}"
  )

# Function to add Go projects
function(add_go_project PROJECT_NAME SOURCE_DIR DEPENDENCIES)
  add_custom_command(
    OUTPUT ${CMAKE_BINARY_DIR}/${PROJECT_NAME}
    COMMAND CGO_ENABLED=1 GOOS="${CMAKE_SYSTEM_NAME}" GOARCH="${CMAKE_SYSTEM_PROCESSOR}" GOARM="" go build -buildmode=c-shared -o ${CMAKE_BINARY_DIR}/${PROJECT_NAME} ${SOURCE_DIR}
    DEPENDS ${SOURCE_DIR}/*.go ${DEPENDENCIES}
    WORKING_DIRECTORY ${CMAKE_SOURCE_DIR}
    COMMENT "Building Go project ${PROJECT_NAME}"
  )
  add_custom_target(${PROJECT_NAME} ALL DEPENDS ${CMAKE_BINARY_DIR}/${PROJECT_NAME})
  add_dependencies(${PROJECT_NAME} ${DEPENDENCIES})
endfunction()

# Add the Go projects
# add_go_project(xyz_plus ${CMAKE_SOURCE_DIR}/xyz_plus "")
add_go_project(scanScheduler ${CMAKE_SOURCE_DIR}/scanScheduler xyz_plus)
add_go_project(pron ${CMAKE_SOURCE_DIR}/pron scanScheduler)
add_go_project(flutter_pron ${CMAKE_SOURCE_DIR}/flutter_pron pron)
