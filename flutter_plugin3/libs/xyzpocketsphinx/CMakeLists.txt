configure_file(include/config.h.in include/config.h)
add_definitions(-DHAVE_CONFIG_H)


add_library(xyzpocketsphinx SHARED
src/libpocketsphinx/acmod.c
src/libpocketsphinx/allphone_search.c
src/libpocketsphinx/bin_mdef.c
src/libpocketsphinx/blkarray_list.c
src/libpocketsphinx/dict.c
src/libpocketsphinx/dict2pid.c
src/libpocketsphinx/fsg_history.c
src/libpocketsphinx/fsg_lextree.c
src/libpocketsphinx/fsg_search.c
src/libpocketsphinx/hmm.c
src/libpocketsphinx/kws_detections.c
src/libpocketsphinx/kws_search.c
src/libpocketsphinx/mdef.c
src/libpocketsphinx/ms_gauden.c
src/libpocketsphinx/ms_mgau.c
src/libpocketsphinx/ms_senone.c
src/libpocketsphinx/ngram_search.c
src/libpocketsphinx/ngram_search_fwdflat.c
src/libpocketsphinx/ngram_search_fwdtree.c
src/libpocketsphinx/phone_loop_search.c
src/libpocketsphinx/pocketsphinx.c
src/libpocketsphinx/ps_alignment.c
src/libpocketsphinx/ps_lattice.c
src/libpocketsphinx/ps_mllr.c
src/libpocketsphinx/ptm_mgau.c
src/libpocketsphinx/s2_semi_mgau.c
src/libpocketsphinx/state_align_search.c
src/libpocketsphinx/tmat.c
src/libpocketsphinx/vector.c
)
set_property(TARGET xyzpocketsphinx PROPERTY C_STANDARD 11)

target_include_directories(
xyzpocketsphinx PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/include
xyzpocketsphinx PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/src/libpocketsphinx

# xyzpocketsphinx PRIVATE ${CMAKE_BINARY_DIR}
# xyzpocketsphinx PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}
# xyzpocketsphinx PUBLIC ${CMAKE_SOURCE_DIR}/include
# xyzpocketsphinx PUBLIC ${CMAKE_BINARY_DIR}/include
# xyzpocketsphinx INTERFACE ${CMAKE_SOURCE_DIR}/include
# xyzpocketsphinx INTERFACE ${CMAKE_BINARY_DIR}/include
)