#!/bin/sh
set -e
if test "$CONFIGURATION" = "Debug"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/Debug/libps_plus.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/Debug/libps_plus.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/Debug/libps_plus.dylib
fi
if test "$CONFIGURATION" = "Release"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/Release/libps_plus.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/Release/libps_plus.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/Release/libps_plus.dylib
fi
if test "$CONFIGURATION" = "MinSizeRel"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/MinSizeRel/libps_plus.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/MinSizeRel/libps_plus.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/MinSizeRel/libps_plus.dylib
fi
if test "$CONFIGURATION" = "RelWithDebInfo"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/RelWithDebInfo/libps_plus.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/RelWithDebInfo/libps_plus.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/RelWithDebInfo/libps_plus.dylib
fi

