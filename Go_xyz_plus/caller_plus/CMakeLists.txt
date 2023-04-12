# GOOS="${CMAKE_SYSTEM_NAME}" GOARCH="${CMAKE_SYSTEM_PROCESSOR}"
message(STATUS "ps_plus lib location: ${PROJECT_BINARY_DIR}")
message(STATUS "GOOS: ${CMAKE_SYSTEM_NAME}")
message(STATUS "GOARCH: ${CMAKE_SYSTEM_PROCESSOR}")
# Compile the Go executable
add_custom_command(
  OUTPUT "${CMAKE_BINARY_DIR}/exec"
  COMMAND CGO_LDFLAGS="-L${PROJECT_BINARY_DIR} -lps_plus -Wl,-rpath,${PROJECT_BINARY_DIR}" go build -o caller_plus_exec main.go data.go strings.go
#   DEPENDS ${CMAKE_SOURCE_DIR}/*.go 
#   WORKING_DIRECTORY ${CMAKE_SOURCE_DIR}
#   COMMENT "Building Go executable ${GO_EXECUTABLE_NAME}"
)
# add_custom_target(${CGO_MODULE_NAME} ALL DEPENDS ${CMAKE_BINARY_DIR}/${CGO_MODULE_NAME})
# add_custom_target(${GO_EXECUTABLE_NAME} ALL DEPENDS ${CMAKE_BINARY_DIR})
# # add_dependencies(${CGO_MODULE_NAME} ${CPP_LIB_NAME})
# add_dependencies(${GO_EXECUTABLE_NAME} ${CGO_MODULE_NAME})
