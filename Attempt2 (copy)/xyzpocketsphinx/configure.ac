dnl Welcome to the Sphinx automated build system.
dnl try not to hurt yourself ;)

AC_INIT(xyzpocketsphinx, 5prealpha)
AC_CONFIG_HEADERS([include/config.h])
AM_INIT_AUTOMAKE([no-define foreign])
AC_CONFIG_MACRO_DIR([m4])

CFLAGS=${CFLAGS:--g -O2 -Wall}

AC_CANONICAL_HOST
AC_PROG_CC
AC_CHECK_TYPES(long long)
AC_CHECK_SIZEOF(long long)

LT_INIT

dnl
dnl Check for pkgconfig
dnl
PKG_PROG_PKG_CONFIG

dnl
dnl Check for Doxygen, and build dox if present
dnl
AC_CHECK_PROG(HAVE_DOXYGEN, doxygen, yes, no)
AM_CONDITIONAL(BUILD_DOXYGEN, test "x$HAVE_DOXYGEN" = "xyes")

dnl Check for SWIG and Python
AC_ARG_WITH(python,
	AS_HELP_STRING([--with-python],
                 [Enable Python extension, built with swig, enabled by default]))

if test "x${with_python}" != "xno"; then
  AM_PATH_PYTHON(2.6, [], [AC_MSG_ERROR([python not found])])
  AX_PYTHON_DEVEL()
  AX_PKG_SWIG(2.0, [], [AC_MSG_ERROR([swig not found])])
fi

AC_SUBST([PYTHON_CPPFLAGS])
AM_CONDITIONAL(BUILD_SWIG, test "x$SWIG" != "x")

dnl swig python check

dnl
dnl Now check for GStreamer, and build the plugin if it's available
dnl
GST_MAJORMINOR=1.0
PKG_CHECK_MODULES(GStreamer, [gstreamer-$GST_MAJORMINOR >= 1.0
		       gstreamer-base-$GST_MAJORMINOR >= 1.0
		       gstreamer-plugins-base-$GST_MAJORMINOR >= 1.0],
                  HAVE_GST=yes, HAVE_GST=no)
dnl Don't build GStreamer when cross-compiling
AM_CONDITIONAL(BUILD_GST, test x$cross_compiling != xyes && test "x$HAVE_GST" = "xyes")
GST_CFLAGS="$GStreamer_CFLAGS $GStreamer_ERROR"
GST_LIBS="$GStreamer_LIBS"
AC_SUBST(GST_MAJORMINOR)
AC_SUBST(GST_CFLAGS)
AC_SUBST(GST_LIBS)

dnl set the plugindir where plugins should be installed
if test "x${prefix}" = "x$HOME"; then
  plugindir="$HOME/.gstreamer-$GST_MAJORMINOR/plugins"
else
  plugindir="\$(libdir)/gstreamer-$GST_MAJORMINOR"
fi
AC_SUBST(plugindir)

dnl set proper LDFLAGS for plugins
GST_PLUGIN_LDFLAGS='-module -avoid-version -export-symbols-regex [_]*\(gst_\|Gst\|GST_\).*'
AC_SUBST(GST_PLUGIN_LDFLAGS)

dnl
dnl Get xyzSphinxBase from command line if given
dnl
AC_ARG_WITH(xyzsphinxbase,
	AS_HELP_STRING([--with-xyzxyzsphinxbase=DIRECTORY],
			[Look for xyzSphinxBase installation in DIRECTORY.  If this is 'auto',
			 the system-wide installation will be used.]),
			 xyzxyzsphinxbase=$withval)


dnl
dnl Check for xyzSphinxBase in parent directories
dnl
if test x$xyzsphinxbase = x; then
   dn=`dirname $0`
   case "$dn" in
	.)
	sbdir="`pwd`/.."
	;;
	[\\/]* | ?:[\\/]*)
	sbdir="$dn/.."
	;;
	*)
	sbdir="`pwd`/$dn/.."
	;;
   esac
   # Look for xyzsphinxbase in the parent directory
   for sb in "$sbdir/xyzsphinxbase" \
	  "$sbdir/xyzsphinxbase"*; do
       AC_MSG_CHECKING([for xyzsphinxbase in $sb])
       if test -f "$sb/src/libsphinxbase/libxyzsphinxbase.la"; then
          xyzsphinxbase="$sb"
	  AC_MSG_RESULT(yes)
	  break
       else
	  AC_MSG_RESULT(no)
       fi
   done
fi

dnl
dnl Check for system xyzxyzSphinxBase if none was passed to us
dnl Also allow --with-xyzsphinxbase=auto to use system one explicitly
dnl
if test x$xyzsphinxbase = x || test x$xyzsphinxbase = xauto; then
   xyzsphinxbase=

   if test "x$PKG_CONFIG" = "x"; then
      XYZSPHINXBASE_CFLAGS = "-I/usr/include/xyzsphinxbase -I/usr/local/include/xyzsphinxbase"
      XYZSPHINXBASE_LIBS = "-lxyzsphinxbase"
      XYZSPHINXBASE_PREFIX="/usr/local"
   else
      PKG_CHECK_MODULES(XYZSPHINXBASE, [xyzsphinxbase],,[
      AC_MSG_FAILURE(dnl
[xyzSphinxBase was not found on your system.
Make sure that you have installed it and that the
PKG_CONFIG_PATH environment variable is set correctly, if
it was installed in a non-standard prefix.])])
      XYZSPHINXBASE_PREFIX=`$PKG_CONFIG --variable=prefix xyzsphinxbase`
   fi

   LIBS="$LIBS $XYZSPHINXBASE_LIBS"
   CPPFLAGS="$CPPFLAGS $XYZSPHINXBASE_CFLAGS"
   SPHINXBASE_SWIG="$XYZSPHINXBASE_PREFIX/share/xyzsphinxbase/swig"
   AC_CHECK_HEADER(sphinx_config.h,,[AC_MSG_FAILURE([xyzSphinxBase was not found on your system.])])
else
   LIBS="$LIBS -lxyzsphinxbase"
   LDFLAGS="$LDFLAGS -L$xyzsphinxbase/lib -L$xyzsphinxbase/src/libsphinxad -L$xyzsphinxbase/src/libsphinxbase"
   CPPFLAGS="$CPPFLAGS -I$xyzsphinxbase/include -I$xyzsphinxbase/include/xyzsphinxbase"
   SPHINXBASE_SWIG="$xyzsphinxbase/swig"
fi

AC_SUBST(SPHINXBASE_SWIG)

AC_OUTPUT([
xyzpocketsphinx.pc
Makefile
include/Makefile
src/Makefile
swig/Makefile
swig/python/Makefile
swig/python/test/Makefile
src/libpocketsphinx/Makefile
src/programs/Makefile
src/gst-plugin/Makefile
doc/Makefile
doc/doxyfile
model/Makefile
test/Makefile
test/testfuncs.sh
test/unit/Makefile
test/regression/Makefile
])
