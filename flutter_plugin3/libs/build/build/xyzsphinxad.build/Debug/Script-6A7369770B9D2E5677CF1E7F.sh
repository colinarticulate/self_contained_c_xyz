#!/bin/sh
set -e
if test "$CONFIGURATION" = "Debug"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/Debug/libxyzsphinxad.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/Debug/libxyzsphinxad.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/Debug/libxyzsphinxad.dylib
fi
if test "$CONFIGURATION" = "Release"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/Release/libxyzsphinxad.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/Release/libxyzsphinxad.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/Release/libxyzsphinxad.dylib
fi
if test "$CONFIGURATION" = "MinSizeRel"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/MinSizeRel/libxyzsphinxad.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/MinSizeRel/libxyzsphinxad.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/MinSizeRel/libxyzsphinxad.dylib
fi
if test "$CONFIGURATION" = "RelWithDebInfo"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/RelWithDebInfo/libxyzsphinxad.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/RelWithDebInfo/libxyzsphinxad.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzsphinxbase/RelWithDebInfo/libxyzsphinxad.dylib
fi

