bin_PROGRAMS = \
	xyzpocketsphinx_batch \
	xyzpocketsphinx_continuous \
	xyzpocketsphinx_mdef_convert

xyzpocketsphinx_mdef_convert_SOURCES = mdef_convert.c
xyzpocketsphinx_mdef_convert_LDADD = \
	$(top_builddir)/src/libpocketsphinx/libxyzpocketsphinx.la

xyzpocketsphinx_batch_SOURCES = batch.c
xyzpocketsphinx_batch_LDADD = \
	$(top_builddir)/src/libpocketsphinx/libxyzpocketsphinx.la

xyzpocketsphinx_continuous_SOURCES = continuous.c
xyzpocketsphinx_continuous_LDADD = \
	$(top_builddir)/src/libpocketsphinx/libxyzpocketsphinx.la -lxyzsphinxad

AM_CFLAGS =-I$(top_srcdir)/include \
	-I$(top_srcdir)/src/libpocketsphinx \
	-I$(top_builddir)/include
