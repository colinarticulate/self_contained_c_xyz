#!/bin/sh
set -e
if test "$CONFIGURATION" = "Debug"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs
  make -f /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs/CMakeScripts/ReRunCMake.make
fi
if test "$CONFIGURATION" = "Release"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs
  make -f /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs/CMakeScripts/ReRunCMake.make
fi
if test "$CONFIGURATION" = "MinSizeRel"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs
  make -f /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs/CMakeScripts/ReRunCMake.make
fi
if test "$CONFIGURATION" = "RelWithDebInfo"; then :
  cd /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs
  make -f /Users/davidbarbera/Repositories/self_contained_c_xyz/flutter_plugin2/libs/CMakeScripts/ReRunCMake.make
fi
