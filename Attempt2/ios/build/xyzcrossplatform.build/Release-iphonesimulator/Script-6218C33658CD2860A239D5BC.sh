#!/bin/sh
set -e
if test "$CONFIGURATION" = "Debug"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzcrossplatform
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzcrossplatform/Debug${EFFECTIVE_PLATFORM_NAME}/libxyzcrossplatform.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzcrossplatform/Debug${EFFECTIVE_PLATFORM_NAME}/libxyzcrossplatform.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzcrossplatform/Debug${EFFECTIVE_PLATFORM_NAME}/libxyzcrossplatform.dylib
fi
if test "$CONFIGURATION" = "Release"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzcrossplatform
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzcrossplatform/Release${EFFECTIVE_PLATFORM_NAME}/libxyzcrossplatform.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzcrossplatform/Release${EFFECTIVE_PLATFORM_NAME}/libxyzcrossplatform.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzcrossplatform/Release${EFFECTIVE_PLATFORM_NAME}/libxyzcrossplatform.dylib
fi
if test "$CONFIGURATION" = "MinSizeRel"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzcrossplatform
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzcrossplatform/MinSizeRel${EFFECTIVE_PLATFORM_NAME}/libxyzcrossplatform.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzcrossplatform/MinSizeRel${EFFECTIVE_PLATFORM_NAME}/libxyzcrossplatform.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzcrossplatform/MinSizeRel${EFFECTIVE_PLATFORM_NAME}/libxyzcrossplatform.dylib
fi
if test "$CONFIGURATION" = "RelWithDebInfo"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzcrossplatform
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzcrossplatform/RelWithDebInfo${EFFECTIVE_PLATFORM_NAME}/libxyzcrossplatform.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzcrossplatform/RelWithDebInfo${EFFECTIVE_PLATFORM_NAME}/libxyzcrossplatform.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzcrossplatform/RelWithDebInfo${EFFECTIVE_PLATFORM_NAME}/libxyzcrossplatform.dylib
fi

