configure_file(include/config.h.in include/config.h)
configure_file(include/sphinx_config.h.in include/sphinx_config.h)
add_definitions(-DHAVE_CONFIG_H)

add_library(xyzsphinxbase SHARED
    src/libsphinxbase/fe/fe_interface.c
    src/libsphinxbase/fe/fe_noise.c
    src/libsphinxbase/fe/fe_prespch_buf.c
    src/libsphinxbase/fe/fe_sigproc.c
    src/libsphinxbase/fe/fe_warp.c
    src/libsphinxbase/fe/fe_warp_affine.c
    src/libsphinxbase/fe/fe_warp_inverse_linear.c
    src/libsphinxbase/fe/fe_warp_piecewise_linear.c
    src/libsphinxbase/fe/fixlog.c
    src/libsphinxbase/fe/yin.c
    src/libsphinxbase/feat/agc.c
    src/libsphinxbase/feat/cmn.c
    src/libsphinxbase/feat/cmn_live.c
    src/libsphinxbase/feat/feat.c
    src/libsphinxbase/feat/lda.c
    src/libsphinxbase/lm/fsg_model.c
    src/libsphinxbase/lm/jsgf.c
    src/libsphinxbase/lm/jsgf_parser.c
    src/libsphinxbase/lm/jsgf_parser.y
    src/libsphinxbase/lm/_jsgf_scanner.l
    src/libsphinxbase/lm/jsgf_scanner.c
    src/libsphinxbase/lm/lm_trie.c
    src/libsphinxbase/lm/lm_trie_quant.c
    src/libsphinxbase/lm/ngram_model.c
    src/libsphinxbase/lm/ngram_model_set.c
    src/libsphinxbase/lm/ngram_model_trie.c
    src/libsphinxbase/lm/ngrams_raw.c
    src/libsphinxbase/util/bio.c
    src/libsphinxbase/util/bitarr.c
    src/libsphinxbase/util/bitvec.c
    src/libsphinxbase/util/blas_lite.c
    src/libsphinxbase/util/case.c
    src/libsphinxbase/util/ckd_alloc.c
    src/libsphinxbase/util/cmd_ln.c
    src/libsphinxbase/util/dtoa.c
    src/libsphinxbase/util/err.c
    src/libsphinxbase/util/errno.c
    src/libsphinxbase/util/f2c_lite.c
    src/libsphinxbase/util/filename.c
    src/libsphinxbase/util/genrand.c
    src/libsphinxbase/util/glist.c
    src/libsphinxbase/util/hash_table.c
    src/libsphinxbase/util/heap.c
    src/libsphinxbase/util/listelem_alloc.c
    src/libsphinxbase/util/logmath.c
    src/libsphinxbase/util/matrix.c
    src/libsphinxbase/util/mmio.c
    src/libsphinxbase/util/pio.c
    src/libsphinxbase/util/priority_queue.c
    src/libsphinxbase/util/profile.c
    src/libsphinxbase/util/sbthread.c
    src/libsphinxbase/util/slamch.c
    src/libsphinxbase/util/slapack_lite.c
    src/libsphinxbase/util/strfuncs.c
)
set_property(TARGET xyzsphinxbase PROPERTY C_STANDARD 11)


add_library(xyzsphinxad SHARED
    # src/libsphinxad/ad_alsa.c
    src/libsphinxad/ad_base.c
    # src/libsphinxad/ad_openal.c
    # src/libsphinxad/ad_oss.c
    # src/libsphinxad/ad_pulse.c
    # src/libsphinxad/ad_win32.c
)
set_property(TARGET xyzsphinxad PROPERTY C_STANDARD 11)

target_include_directories(
    # xyzsphinxbase PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/include
    # xyzsphinxbase INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}/include
    # xyzsphinxbase INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}/src/libsphinxbase/fe
    # xyzsphinxbase INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}/src/libsphinxbase/feat
    # xyzsphinxbase INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}/src/libsphinxbase/lm
    # xyzsphinxbase INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}/src/libsphinxbase/util
    # xyzsphinxad PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/include
    # xyzsphinxad INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}/include
    # xyzsphinxad INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}/src/libsphinxad

    # xyzsphinxbase PRIVATE ${CMAKE_BINARY_DIR}
    # xyzsphinxbase PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}
    # xyzsphinxbase PUBLIC ${CMAKE_SOURCE_DIR}/include
    # xyzsphinxbase PUBLIC ${CMAKE_BINARY_DIR}/include
    xyzsphinxbase PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/include
    # xyzsphinxbase INTERFACE ${CMAKE_SOURCE_DIR}/include
    # xyzsphinxbase INTERFACE ${CMAKE_BINARY_DIR}/include
    xyzsphinxbase PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/src/libsphinxbase/fe
    xyzsphinxbase PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/src/libsphinxbase/feat
    xyzsphinxbase PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/src/libsphinxbase/lm
    xyzsphinxbase PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/src/libsphinxbase/util

    # xyzsphinxad PRIVATE ${CMAKE_BINARY_DIR}
    # xyzsphinxad PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}
    xyzsphinxad INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}/include/xyzsphinxbase
    xyzsphinxad PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/include
    
    # xyzsphinxad PUBLIC ${CMAKE_BINARY_DIR}/include
    # xyzsphinxad INTERFACE ${CMAKE_SOURCE_DIR}/include
    # xyzsphinxad INTERFACE ${CMAKE_BINARY_DIR}/include
    #xyzsphinxad PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}/src/libsphinxad
)
