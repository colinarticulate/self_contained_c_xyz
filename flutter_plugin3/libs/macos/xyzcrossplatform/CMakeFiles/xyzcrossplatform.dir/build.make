# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.26

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
CMAKE_COMMAND = /Applications/CMake.app/Contents/bin/cmake

# The command to remove a file.
RM = /Applications/CMake.app/Contents/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/macos

# Include any dependencies generated for this target.
include xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/depend.make
# Include any dependencies generated by the compiler for this target.
include xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/compiler_depend.make

# Include the progress variables for this target.
include xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/progress.make

# Include the compile flags for this target's objects.
include xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/flags.make

xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o: xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/flags.make
xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o: /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/xyzcrossplatform/crossplatform.c
xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o: xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/macos/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o"
	cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/macos/xyzcrossplatform && /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o -MF CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o.d -o CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o -c /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/xyzcrossplatform/crossplatform.c

xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/crossplatform.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/xyzcrossplatform.dir/crossplatform.c.i"
	cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/macos/xyzcrossplatform && /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/xyzcrossplatform/crossplatform.c > CMakeFiles/xyzcrossplatform.dir/crossplatform.c.i

xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/crossplatform.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/xyzcrossplatform.dir/crossplatform.c.s"
	cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/macos/xyzcrossplatform && /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/xyzcrossplatform/crossplatform.c -o CMakeFiles/xyzcrossplatform.dir/crossplatform.c.s

# Object files for target xyzcrossplatform
xyzcrossplatform_OBJECTS = \
"CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o"

# External object files for target xyzcrossplatform
xyzcrossplatform_EXTERNAL_OBJECTS =

xyzcrossplatform/libxyzcrossplatform.dylib: xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/crossplatform.c.o
xyzcrossplatform/libxyzcrossplatform.dylib: xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/build.make
xyzcrossplatform/libxyzcrossplatform.dylib: xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/macos/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking C shared library libxyzcrossplatform.dylib"
	cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/macos/xyzcrossplatform && $(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/xyzcrossplatform.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/build: xyzcrossplatform/libxyzcrossplatform.dylib
.PHONY : xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/build

xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/clean:
	cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/macos/xyzcrossplatform && $(CMAKE_COMMAND) -P CMakeFiles/xyzcrossplatform.dir/cmake_clean.cmake
.PHONY : xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/clean

xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/depend:
	cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/macos && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/xyzcrossplatform /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/macos /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/macos/xyzcrossplatform /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/macos/xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : xyzcrossplatform/CMakeFiles/xyzcrossplatform.dir/depend
