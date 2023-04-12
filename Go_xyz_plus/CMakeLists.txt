cmake_minimum_required(VERSION 3.15)
project(GoCgoCppExample C CXX)# Go)

# Set the C++ library, CGO module, and Go executable names
set(CPP_LIB_NAME ps_plus)
set(CGO_MODULE_NAME xyz_plus)
set(GO_EXECUTABLE_NAME caller_plus)

# Add the C++ library
add_subdirectory(ps_plus)
add_subdirectory(caller_plus)

    