bin_PROGRAMS = xyzsphinx_fe

xyzsphinx_fe_SOURCES = sphinx_fe.c

noinst_HEADERS = cmd_ln_defn.h \
	sphinx_wave2feat.h

xyzsphinx_fe_LDADD = \
	$(top_builddir)/src/libsphinxbase/libxyzsphinxbase.la

AM_CFLAGS =-I$(top_srcdir)/include \
	-I$(top_builddir)/include \
	-I$(top_srcdir)/src/libsphinxbase/fe

