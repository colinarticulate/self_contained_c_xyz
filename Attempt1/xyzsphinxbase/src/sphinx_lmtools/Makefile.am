bin_PROGRAMS = xyzsphinx_lm_eval xyzsphinx_lm_convert

xyzsphinx_lm_eval_SOURCES = sphinx_lm_eval.c
xyzsphinx_lm_eval_LDADD = \
	$(top_builddir)/src/libsphinxbase/libxyzsphinxbase.la \
	$(top_builddir)/src/libsphinxbase/lm/libxyzsphinxlm.la

xyzsphinx_lm_convert_SOURCES = sphinx_lm_convert.c
xyzsphinx_lm_convert_LDADD = \
	$(top_builddir)/src/libsphinxbase/libxyzsphinxbase.la \
	$(top_builddir)/src/libsphinxbase/lm/libxyzsphinxlm.la

AM_CFLAGS =-I$(top_srcdir)/include -I$(top_builddir)/include \
	-I$(top_srcdir)/src/libsphinxlm
