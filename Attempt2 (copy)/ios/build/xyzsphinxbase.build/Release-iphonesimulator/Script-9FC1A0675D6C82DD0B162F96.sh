#!/bin/sh
set -e
if test "$CONFIGURATION" = "Debug"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzsphinxbase
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzsphinxbase/Debug${EFFECTIVE_PLATFORM_NAME}/libxyzsphinxbase.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzsphinxbase/Debug${EFFECTIVE_PLATFORM_NAME}/libxyzsphinxbase.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzsphinxbase/Debug${EFFECTIVE_PLATFORM_NAME}/libxyzsphinxbase.dylib
fi
if test "$CONFIGURATION" = "Release"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzsphinxbase
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzsphinxbase/Release${EFFECTIVE_PLATFORM_NAME}/libxyzsphinxbase.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzsphinxbase/Release${EFFECTIVE_PLATFORM_NAME}/libxyzsphinxbase.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzsphinxbase/Release${EFFECTIVE_PLATFORM_NAME}/libxyzsphinxbase.dylib
fi
if test "$CONFIGURATION" = "MinSizeRel"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzsphinxbase
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzsphinxbase/MinSizeRel${EFFECTIVE_PLATFORM_NAME}/libxyzsphinxbase.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzsphinxbase/MinSizeRel${EFFECTIVE_PLATFORM_NAME}/libxyzsphinxbase.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzsphinxbase/MinSizeRel${EFFECTIVE_PLATFORM_NAME}/libxyzsphinxbase.dylib
fi
if test "$CONFIGURATION" = "RelWithDebInfo"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzsphinxbase
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzsphinxbase/RelWithDebInfo${EFFECTIVE_PLATFORM_NAME}/libxyzsphinxbase.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzsphinxbase/RelWithDebInfo${EFFECTIVE_PLATFORM_NAME}/libxyzsphinxbase.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzsphinxbase/RelWithDebInfo${EFFECTIVE_PLATFORM_NAME}/libxyzsphinxbase.dylib
fi

