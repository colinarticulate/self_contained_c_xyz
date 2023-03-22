#!/bin/sh
set -e
if test "$CONFIGURATION" = "Debug"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/Debug${EFFECTIVE_PLATFORM_NAME}/libps_plus.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/Debug${EFFECTIVE_PLATFORM_NAME}/libps_plus.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/Debug${EFFECTIVE_PLATFORM_NAME}/libps_plus.dylib
fi
if test "$CONFIGURATION" = "Release"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/Release${EFFECTIVE_PLATFORM_NAME}/libps_plus.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/Release${EFFECTIVE_PLATFORM_NAME}/libps_plus.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/Release${EFFECTIVE_PLATFORM_NAME}/libps_plus.dylib
fi
if test "$CONFIGURATION" = "MinSizeRel"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/MinSizeRel${EFFECTIVE_PLATFORM_NAME}/libps_plus.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/MinSizeRel${EFFECTIVE_PLATFORM_NAME}/libps_plus.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/MinSizeRel${EFFECTIVE_PLATFORM_NAME}/libps_plus.dylib
fi
if test "$CONFIGURATION" = "RelWithDebInfo"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/RelWithDebInfo${EFFECTIVE_PLATFORM_NAME}/libps_plus.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/RelWithDebInfo${EFFECTIVE_PLATFORM_NAME}/libps_plus.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/RelWithDebInfo${EFFECTIVE_PLATFORM_NAME}/libps_plus.dylib
fi

