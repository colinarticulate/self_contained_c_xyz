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
CMAKE_BINARY_DIR = /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/build

# Include any dependencies generated for this target.
include CMakeFiles/ps_plus.dir/depend.make
# Include any dependencies generated by the compiler for this target.
include CMakeFiles/ps_plus.dir/compiler_depend.make

# Include the progress variables for this target.
include CMakeFiles/ps_plus.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/ps_plus.dir/flags.make

CMakeFiles/ps_plus.dir/c_demo.cpp.o: CMakeFiles/ps_plus.dir/flags.make
CMakeFiles/ps_plus.dir/c_demo.cpp.o: ../c_demo.cpp
CMakeFiles/ps_plus.dir/c_demo.cpp.o: CMakeFiles/ps_plus.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/ps_plus.dir/c_demo.cpp.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/ps_plus.dir/c_demo.cpp.o -MF CMakeFiles/ps_plus.dir/c_demo.cpp.o.d -o CMakeFiles/ps_plus.dir/c_demo.cpp.o -c /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/c_demo.cpp

CMakeFiles/ps_plus.dir/c_demo.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/ps_plus.dir/c_demo.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/c_demo.cpp > CMakeFiles/ps_plus.dir/c_demo.cpp.i

CMakeFiles/ps_plus.dir/c_demo.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/ps_plus.dir/c_demo.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/c_demo.cpp -o CMakeFiles/ps_plus.dir/c_demo.cpp.s

# Object files for target ps_plus
ps_plus_OBJECTS = \
"CMakeFiles/ps_plus.dir/c_demo.cpp.o"

# External object files for target ps_plus
ps_plus_EXTERNAL_OBJECTS =

../bin/ps_plus: CMakeFiles/ps_plus.dir/c_demo.cpp.o
../bin/ps_plus: CMakeFiles/ps_plus.dir/build.make
../bin/ps_plus: xyzpocketsphinx/libxyzpocketsphinx.so
../bin/ps_plus: xyzsphinxbase/libxyzsphinxad.so
../bin/ps_plus: xyzsphinxbase/libxyzsphinxbase.so
../bin/ps_plus: xyzcrossplatform/libxyzcrossplatform.so
../bin/ps_plus: CMakeFiles/ps_plus.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking CXX executable ../bin/ps_plus"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/ps_plus.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/ps_plus.dir/build: ../bin/ps_plus
.PHONY : CMakeFiles/ps_plus.dir/build

CMakeFiles/ps_plus.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/ps_plus.dir/cmake_clean.cmake
.PHONY : CMakeFiles/ps_plus.dir/clean

CMakeFiles/ps_plus.dir/depend:
	cd /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/build && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2 /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2 /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/build /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/build /home/dbarbera/Repositories/self_contained_c_xyz/Attempt2/build/CMakeFiles/ps_plus.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/ps_plus.dir/depend

