#!/bin/sh
set -e
if test "$CONFIGURATION" = "Debug"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios
  make -f /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/CMakeScripts/ReRunCMake.make
fi
if test "$CONFIGURATION" = "Release"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios
  make -f /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/CMakeScripts/ReRunCMake.make
fi
if test "$CONFIGURATION" = "MinSizeRel"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios
  make -f /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/CMakeScripts/ReRunCMake.make
fi
if test "$CONFIGURATION" = "RelWithDebInfo"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios
  make -f /Users/davidbarbera/Repositories/self_contained_c_xyz/Attempt2/ios/CMakeScripts/ReRunCMake.make
fi

