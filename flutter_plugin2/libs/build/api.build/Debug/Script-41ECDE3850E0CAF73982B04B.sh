#!/bin/sh
set -e
if test "$CONFIGURATION" = "Debug"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs/Debug/libapi.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs/Debug/libapi.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs/Debug/libapi.dylib
fi
if test "$CONFIGURATION" = "Release"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs/Release/libapi.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs/Release/libapi.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs/Release/libapi.dylib
fi
if test "$CONFIGURATION" = "MinSizeRel"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs/MinSizeRel/libapi.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs/MinSizeRel/libapi.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs/MinSizeRel/libapi.dylib
fi
if test "$CONFIGURATION" = "RelWithDebInfo"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs/RelWithDebInfo/libapi.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs/RelWithDebInfo/libapi.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs/RelWithDebInfo/libapi.dylib
fi

