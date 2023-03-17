#!/bin/sh
set -e
if test "$CONFIGURATION" = "Debug"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzpocketsphinx
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzpocketsphinx/Debug/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzpocketsphinx/Debug/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzpocketsphinx/Debug/libxyzpocketsphinx.dylib
fi
if test "$CONFIGURATION" = "Release"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzpocketsphinx
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzpocketsphinx/Release/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzpocketsphinx/Release/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzpocketsphinx/Release/libxyzpocketsphinx.dylib
fi
if test "$CONFIGURATION" = "MinSizeRel"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzpocketsphinx
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzpocketsphinx/MinSizeRel/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzpocketsphinx/MinSizeRel/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzpocketsphinx/MinSizeRel/libxyzpocketsphinx.dylib
fi
if test "$CONFIGURATION" = "RelWithDebInfo"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzpocketsphinx
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzpocketsphinx/RelWithDebInfo/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzpocketsphinx/RelWithDebInfo/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin3/libs/build/xyzpocketsphinx/RelWithDebInfo/libxyzpocketsphinx.dylib
fi

