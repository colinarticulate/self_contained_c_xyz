#!/bin/sh
set -e
if test "$CONFIGURATION" = "Debug"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/Debug/libxyzsphinxbase.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/Debug/libxyzsphinxbase.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/Debug/libxyzsphinxbase.dylib
fi
if test "$CONFIGURATION" = "Release"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/Release/libxyzsphinxbase.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/Release/libxyzsphinxbase.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/Release/libxyzsphinxbase.dylib
fi
if test "$CONFIGURATION" = "MinSizeRel"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/MinSizeRel/libxyzsphinxbase.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/MinSizeRel/libxyzsphinxbase.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/MinSizeRel/libxyzsphinxbase.dylib
fi
if test "$CONFIGURATION" = "RelWithDebInfo"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/RelWithDebInfo/libxyzsphinxbase.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/RelWithDebInfo/libxyzsphinxbase.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/RelWithDebInfo/libxyzsphinxbase.dylib
fi

