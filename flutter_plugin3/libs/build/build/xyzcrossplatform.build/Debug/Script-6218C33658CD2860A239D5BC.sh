#!/bin/sh
set -e
if test "$CONFIGURATION" = "Debug"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzcrossplatform
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzcrossplatform/Debug/libxyzcrossplatform.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzcrossplatform/Debug/libxyzcrossplatform.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzcrossplatform/Debug/libxyzcrossplatform.dylib
fi
if test "$CONFIGURATION" = "Release"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzcrossplatform
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzcrossplatform/Release/libxyzcrossplatform.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzcrossplatform/Release/libxyzcrossplatform.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzcrossplatform/Release/libxyzcrossplatform.dylib
fi
if test "$CONFIGURATION" = "MinSizeRel"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzcrossplatform
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzcrossplatform/MinSizeRel/libxyzcrossplatform.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzcrossplatform/MinSizeRel/libxyzcrossplatform.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzcrossplatform/MinSizeRel/libxyzcrossplatform.dylib
fi
if test "$CONFIGURATION" = "RelWithDebInfo"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzcrossplatform
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzcrossplatform/RelWithDebInfo/libxyzcrossplatform.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzcrossplatform/RelWithDebInfo/libxyzcrossplatform.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzcrossplatform/RelWithDebInfo/libxyzcrossplatform.dylib
fi

