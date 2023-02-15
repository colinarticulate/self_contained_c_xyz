# project(xyzpocketsphinx)

# set(libpocketsphinx_sources
# acmod.c
# allphone_search.c
# bin_mdef.c
# blkarray_list.c
# dict.c
# dict2pid.c
# fsg_history.c
# fsg_lextree.c
# fsg_search.c
# hmm.c
# kws_detections.c
# kws_search.c
# mdef.c
# ms_gauden.c
# ms_mgau.c
# ms_senone.c
# ngram_search.c
# ngram_search_fwdflat.c
# ngram_search_fwdtree.c
# phone_loop_search.c
# pocketsphinx.c
# ps_alignment.c
# ps_lattice.c
# ps_mllr.c
# ptm_mgau.c
# s2_semi_mgau.c
# state_align_search.c
# tmat.c
# vector.c
# acmod.h
# allphone_search.h
# bin_mdef.h
# blkarray_list.h
# dict.h
# dict2pid.h
# fsg_history.h
# fsg_lextree.h
# fsg_search_internal.h
# hmm.h
# kws_detections.h
# kws_search.h
# mdef.h
# ms_gauden.h
# ms_mgau.h
# ms_senone.h
# ngram_search.h
# ngram_search_fwdflat.h
# ngram_search_fwdtree.h
# phone_loop_search.h
# pocketsphinx_internal.h
# ps_alignment.h
# ps_lattice_internal.h
# ptm_mgau.h
# s2_semi_mgau.h
# s3types.h
# state_align_search.h
# tied_mgau_common.h
# tmat.h
# vector.h
# )

add_library(libpocketsphinx SHARED $(libpocketsphinx_sources))

target_include_directories(libpocketsphinx
          INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}
          )

target_link_libraries(libpocketsphinx tutorial_compiler_flags)