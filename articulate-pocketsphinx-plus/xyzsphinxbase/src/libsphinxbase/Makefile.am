SUBDIRS = util fe feat lm

lib_LTLIBRARIES = libxyzsphinxbase.la

libxyzsphinxbase_la_SOURCES =

libxyzsphinxbase_la_LDFLAGS = -version-info 3:0:0
libxyzsphinxbase_la_LIBADD = $(LTLIBICONV) \
	util/libxyzsphinxutil.la \
	fe/libxyzsphinxfe.la \
	feat/libxyzsphinxfeat.la \
	lm/libxyzsphinxlm.la
