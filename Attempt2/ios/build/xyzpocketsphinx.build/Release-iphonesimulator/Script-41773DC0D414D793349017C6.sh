#!/bin/sh
set -e
if test "$CONFIGURATION" = "Debug"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzpocketsphinx
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzpocketsphinx/Debug${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzpocketsphinx/Debug${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzpocketsphinx/Debug${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib
fi
if test "$CONFIGURATION" = "Release"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzpocketsphinx
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzpocketsphinx/Release${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzpocketsphinx/Release${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzpocketsphinx/Release${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib
fi
if test "$CONFIGURATION" = "MinSizeRel"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzpocketsphinx
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzpocketsphinx/MinSizeRel${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzpocketsphinx/MinSizeRel${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzpocketsphinx/MinSizeRel${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib
fi
if test "$CONFIGURATION" = "RelWithDebInfo"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzpocketsphinx
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzpocketsphinx/RelWithDebInfo${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzpocketsphinx/RelWithDebInfo${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/xyzpocketsphinx/RelWithDebInfo${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib
fi

