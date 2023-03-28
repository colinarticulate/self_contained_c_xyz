#!/bin/sh
set -e
if test "$CONFIGURATION" = "Debug"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios_build/xyzpocketsphinx
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios_build/xyzpocketsphinx/Debug${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios_build/xyzpocketsphinx/Debug${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios_build/xyzpocketsphinx/Debug${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib
fi
if test "$CONFIGURATION" = "Release"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios_build/xyzpocketsphinx
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios_build/xyzpocketsphinx/Release${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios_build/xyzpocketsphinx/Release${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios_build/xyzpocketsphinx/Release${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib
fi
if test "$CONFIGURATION" = "MinSizeRel"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios_build/xyzpocketsphinx
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios_build/xyzpocketsphinx/MinSizeRel${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios_build/xyzpocketsphinx/MinSizeRel${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios_build/xyzpocketsphinx/MinSizeRel${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib
fi
if test "$CONFIGURATION" = "RelWithDebInfo"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios_build/xyzpocketsphinx
  /Applications/CMake.app/Contents/bin/cmake -E cmake_symlink_library /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios_build/xyzpocketsphinx/RelWithDebInfo${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios_build/xyzpocketsphinx/RelWithDebInfo${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios_build/xyzpocketsphinx/RelWithDebInfo${EFFECTIVE_PLATFORM_NAME}/libxyzpocketsphinx.dylib
fi

