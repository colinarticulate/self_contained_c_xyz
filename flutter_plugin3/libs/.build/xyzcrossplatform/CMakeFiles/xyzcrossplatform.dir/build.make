# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.22

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:

#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:

# Disable VCS-based implicit rules.
% : %,v

# Disable VCS-based implicit rules.
% : RCS/%

# Disable VCS-based implicit rules.
% : RCS/%,v

# Disable VCS-based implicit rules.
% : SCCS/s.%

# Disable VCS-based implicit rules.
% : s.%

.SUFFIXES: .hpux_make_needs_suffix_list

# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

#Suppress display of executed commands.
$(VERBOSE).SILENT:

# A target that is always out of date.
cmake_force:
.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/bin/cmake

# The command to remove a file.
RM = /usr/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/.build

# Include any dependencies generated for this target.
include xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/depend.make
# Include any dependencies generated by the compiler for this target.
include xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/compiler_depend.make

# Include the progress variables for this target.
include xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/progress.make

# Include the compile flags for this target's objects.
include xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/flags.make

xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o: xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/flags.make
xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o: ../xyzcrossplatform/crossplatform.c
xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o: xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/.build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o"
	cd /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/.build/xyzcrossplatform && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o -MF CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o.d -o CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o -c /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/xyzcrossplatform/crossplatform.c

xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/crossplatform.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/xyzcrossplatform.dir/crossplatform.c.i"
	cd /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/.build/xyzcrossplatform && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/xyzcrossplatform/crossplatform.c > CMakeFiles/xyzcrossplatform.dir/crossplatform.c.i

xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/crossplatform.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/xyzcrossplatform.dir/crossplatform.c.s"
	cd /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/.build/xyzcrossplatform && /usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/xyzcrossplatform/crossplatform.c -o CMakeFiles/xyzcrossplatform.dir/crossplatform.c.s

# Object files for target xyzcrossplatform
xyzcrossplatform_OBJECTS = \
"CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o"

# External object files for target xyzcrossplatform
xyzcrossplatform_EXTERNAL_OBJECTS =

xyzcrossplatform/libxyzcrossplatform.so: xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o
xyzcrossplatform/libxyzcrossplatform.so: xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/build.make
xyzcrossplatform/libxyzcrossplatform.so: xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/.build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking C shared library libxyzcrossplatform.so"
	cd /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/.build/xyzcrossplatform && $(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/xyzcrossplatform.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/build: xyzcrossplatform/libxyzcrossplatform.so
.PHONY : xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/build

xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/clean:
	cd /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/.build/xyzcrossplatform && $(CMAKE_COMMAND) -P CMakeFiles/xyzcrossplatform.dir/cmake_clean.cmake
.PHONY : xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/clean

xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/depend:
	cd /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/.build && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2 /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/xyzcrossplatform /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/.build /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/.build/xyzcrossplatform /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/.build/xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/depend

