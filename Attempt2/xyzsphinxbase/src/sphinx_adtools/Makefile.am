bin_PROGRAMS = xyzsphinx_pitch xyzsphinx_cont_seg

xyzsphinx_pitch_SOURCES = sphinx_pitch.c
xyzsphinx_pitch_LDADD = $(top_builddir)/src/libsphinxbase/libxyzsphinxbase.la \
	 $(top_builddir)/src/libsphinxad/libxyzsphinxad.la

xyzsphinx_cont_seg_SOURCES = cont_seg.c
xyzsphinx_cont_seg_LDADD = $(top_builddir)/src/libsphinxbase/libxyzsphinxbase.la \
	 $(top_builddir)/src/libsphinxad/libxyzsphinxad.la

AM_CFLAGS = -I$(top_srcdir)/include -I$(top_builddir)/include
