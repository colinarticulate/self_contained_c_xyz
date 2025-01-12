# c_xyz cmake file

cmake_minimum_required(VERSION 3.4.1)

project(PS_PLUS)

set_source_files_properties(
ps_plus.cpp
xyzsphinxbase/src/libsphinxbase/fe/fe_interface.c
xyzsphinxbase/src/libsphinxbase/fe/fe_noise.c
xyzsphinxbase/src/libsphinxbase/fe/fe_prespch_buf.c
xyzsphinxbase/src/libsphinxbase/fe/fe_sigproc.c
xyzsphinxbase/src/libsphinxbase/fe/fe_warp.c
xyzsphinxbase/src/libsphinxbase/fe/fe_warp_affine.c
xyzsphinxbase/src/libsphinxbase/fe/fe_warp_inverse_linear.c
xyzsphinxbase/src/libsphinxbase/fe/fe_warp_piecewise_linear.c
xyzsphinxbase/src/libsphinxbase/fe/fixlog.c
xyzsphinxbase/src/libsphinxbase/fe/yin.c
xyzsphinxbase/src/libsphinxbase/feat/agc.c
xyzsphinxbase/src/libsphinxbase/feat/cmn.c
xyzsphinxbase/src/libsphinxbase/feat/cmn_live.c
xyzsphinxbase/src/libsphinxbase/feat/feat.c
xyzsphinxbase/src/libsphinxbase/feat/lda.c
xyzsphinxbase/src/libsphinxbase/lm/fsg_model.c
xyzsphinxbase/src/libsphinxbase/lm/jsgf.c
xyzsphinxbase/src/libsphinxbase/lm/jsgf_parser.c
xyzsphinxbase/src/libsphinxbase/lm/jsgf_scanner.c
xyzsphinxbase/src/libsphinxbase/lm/lm_trie.c
xyzsphinxbase/src/libsphinxbase/lm/lm_trie_quant.c
xyzsphinxbase/src/libsphinxbase/lm/ngram_model.c
xyzsphinxbase/src/libsphinxbase/lm/ngram_model_set.c
xyzsphinxbase/src/libsphinxbase/lm/ngram_model_trie.c
xyzsphinxbase/src/libsphinxbase/lm/ngrams_raw.c
xyzsphinxbase/src/libsphinxbase/util/bio.c
xyzsphinxbase/src/libsphinxbase/util/bitarr.c
xyzsphinxbase/src/libsphinxbase/util/bitvec.c
xyzsphinxbase/src/libsphinxbase/util/blas_lite.c
xyzsphinxbase/src/libsphinxbase/util/case.c
xyzsphinxbase/src/libsphinxbase/util/ckd_alloc.c
xyzsphinxbase/src/libsphinxbase/util/cmd_ln.c
xyzsphinxbase/src/libsphinxbase/util/dtoa.c
xyzsphinxbase/src/libsphinxbase/util/err.c
xyzsphinxbase/src/libsphinxbase/util/errno.c
xyzsphinxbase/src/libsphinxbase/util/f2c_lite.c
xyzsphinxbase/src/libsphinxbase/util/filename.c
xyzsphinxbase/src/libsphinxbase/util/genrand.c
xyzsphinxbase/src/libsphinxbase/util/glist.c
xyzsphinxbase/src/libsphinxbase/util/hash_table.c
xyzsphinxbase/src/libsphinxbase/util/heap.c
xyzsphinxbase/src/libsphinxbase/util/listelem_alloc.c
xyzsphinxbase/src/libsphinxbase/util/logmath.c
xyzsphinxbase/src/libsphinxbase/util/matrix.c
xyzsphinxbase/src/libsphinxbase/util/mmio.c
xyzsphinxbase/src/libsphinxbase/util/pio.c
xyzsphinxbase/src/libsphinxbase/util/priority_queue.c
xyzsphinxbase/src/libsphinxbase/util/profile.c
xyzsphinxbase/src/libsphinxbase/util/sbthread.c
xyzsphinxbase/src/libsphinxbase/util/slamch.c
xyzsphinxbase/src/libsphinxbase/util/slapack_lite.c
xyzsphinxbase/src/libsphinxbase/util/strfuncs.c
xyzpocketsphinx/src/libpocketsphinx/acmod.c
xyzpocketsphinx/src/libpocketsphinx/allphone_search.c
xyzpocketsphinx/src/libpocketsphinx/bin_mdef.c
xyzpocketsphinx/src/libpocketsphinx/blkarray_list.c
xyzpocketsphinx/src/libpocketsphinx/dict.c
xyzpocketsphinx/src/libpocketsphinx/dict2pid.c
xyzpocketsphinx/src/libpocketsphinx/fsg_history.c
xyzpocketsphinx/src/libpocketsphinx/fsg_lextree.c
xyzpocketsphinx/src/libpocketsphinx/fsg_search.c
xyzpocketsphinx/src/libpocketsphinx/hmm.c
xyzpocketsphinx/src/libpocketsphinx/kws_detections.c
xyzpocketsphinx/src/libpocketsphinx/kws_search.c
xyzpocketsphinx/src/libpocketsphinx/mdef.c
xyzpocketsphinx/src/libpocketsphinx/ms_gauden.c
xyzpocketsphinx/src/libpocketsphinx/ms_mgau.c
xyzpocketsphinx/src/libpocketsphinx/ms_senone.c
xyzpocketsphinx/src/libpocketsphinx/ngram_search.c
xyzpocketsphinx/src/libpocketsphinx/ngram_search_fwdflat.c
xyzpocketsphinx/src/libpocketsphinx/ngram_search_fwdtree.c
xyzpocketsphinx/src/libpocketsphinx/phone_loop_search.c
xyzpocketsphinx/src/libpocketsphinx/pocketsphinx.c
xyzpocketsphinx/src/libpocketsphinx/ps_alignment.c
xyzpocketsphinx/src/libpocketsphinx/ps_lattice.c
xyzpocketsphinx/src/libpocketsphinx/ps_mllr.c
xyzpocketsphinx/src/libpocketsphinx/ptm_mgau.c
xyzpocketsphinx/src/libpocketsphinx/s2_semi_mgau.c
xyzpocketsphinx/src/libpocketsphinx/state_align_search.c
xyzpocketsphinx/src/libpocketsphinx/tmat.c
xyzpocketsphinx/src/libpocketsphinx/vector.c
PROPERTIES LANGUAGE CXX )





add_compile_options(
"-g"
"-Wall"
"-O0"
"-ggdb"
"-std=c++11"
"-Wno-unused-result"
"-Wno-unused-function"
"-Wno-unused-variable"
"-Wno-unused-const-variable"
"-Wimplicit-function-declaration"
"-Wwritable-strings"
"-lpthread"
"-pthread"
"-lm"
)


add_executable(ps_plus
ps_plus.cpp
xyzsphinxbase/src/libsphinxbase/fe/fe_interface.c
xyzsphinxbase/src/libsphinxbase/fe/fe_noise.c
xyzsphinxbase/src/libsphinxbase/fe/fe_prespch_buf.c
xyzsphinxbase/src/libsphinxbase/fe/fe_sigproc.c
xyzsphinxbase/src/libsphinxbase/fe/fe_warp.c
xyzsphinxbase/src/libsphinxbase/fe/fe_warp_affine.c
xyzsphinxbase/src/libsphinxbase/fe/fe_warp_inverse_linear.c
xyzsphinxbase/src/libsphinxbase/fe/fe_warp_piecewise_linear.c
xyzsphinxbase/src/libsphinxbase/fe/fixlog.c
xyzsphinxbase/src/libsphinxbase/fe/yin.c
xyzsphinxbase/src/libsphinxbase/feat/agc.c
xyzsphinxbase/src/libsphinxbase/feat/cmn.c
xyzsphinxbase/src/libsphinxbase/feat/cmn_live.c
xyzsphinxbase/src/libsphinxbase/feat/feat.c
xyzsphinxbase/src/libsphinxbase/feat/lda.c
xyzsphinxbase/src/libsphinxbase/lm/fsg_model.c
xyzsphinxbase/src/libsphinxbase/lm/jsgf.c
xyzsphinxbase/src/libsphinxbase/lm/jsgf_parser.c
xyzsphinxbase/src/libsphinxbase/lm/jsgf_scanner.c
xyzsphinxbase/src/libsphinxbase/lm/lm_trie.c
xyzsphinxbase/src/libsphinxbase/lm/lm_trie_quant.c
xyzsphinxbase/src/libsphinxbase/lm/ngram_model.c
xyzsphinxbase/src/libsphinxbase/lm/ngram_model_set.c
xyzsphinxbase/src/libsphinxbase/lm/ngram_model_trie.c
xyzsphinxbase/src/libsphinxbase/lm/ngrams_raw.c
xyzsphinxbase/src/libsphinxbase/util/bio.c
xyzsphinxbase/src/libsphinxbase/util/bitarr.c
xyzsphinxbase/src/libsphinxbase/util/bitvec.c
xyzsphinxbase/src/libsphinxbase/util/blas_lite.c
xyzsphinxbase/src/libsphinxbase/util/case.c
xyzsphinxbase/src/libsphinxbase/util/ckd_alloc.c
xyzsphinxbase/src/libsphinxbase/util/cmd_ln.c
xyzsphinxbase/src/libsphinxbase/util/dtoa.c
xyzsphinxbase/src/libsphinxbase/util/err.c
xyzsphinxbase/src/libsphinxbase/util/errno.c
xyzsphinxbase/src/libsphinxbase/util/f2c_lite.c
xyzsphinxbase/src/libsphinxbase/util/filename.c
xyzsphinxbase/src/libsphinxbase/util/genrand.c
xyzsphinxbase/src/libsphinxbase/util/glist.c
xyzsphinxbase/src/libsphinxbase/util/hash_table.c
xyzsphinxbase/src/libsphinxbase/util/heap.c
xyzsphinxbase/src/libsphinxbase/util/listelem_alloc.c
xyzsphinxbase/src/libsphinxbase/util/logmath.c
xyzsphinxbase/src/libsphinxbase/util/matrix.c
xyzsphinxbase/src/libsphinxbase/util/mmio.c
xyzsphinxbase/src/libsphinxbase/util/pio.c
xyzsphinxbase/src/libsphinxbase/util/priority_queue.c
xyzsphinxbase/src/libsphinxbase/util/profile.c
xyzsphinxbase/src/libsphinxbase/util/sbthread.c
xyzsphinxbase/src/libsphinxbase/util/slamch.c
xyzsphinxbase/src/libsphinxbase/util/slapack_lite.c
xyzsphinxbase/src/libsphinxbase/util/strfuncs.c
xyzpocketsphinx/src/libpocketsphinx/acmod.c
xyzpocketsphinx/src/libpocketsphinx/allphone_search.c
xyzpocketsphinx/src/libpocketsphinx/bin_mdef.c
xyzpocketsphinx/src/libpocketsphinx/blkarray_list.c
xyzpocketsphinx/src/libpocketsphinx/dict.c
xyzpocketsphinx/src/libpocketsphinx/dict2pid.c
xyzpocketsphinx/src/libpocketsphinx/fsg_history.c
xyzpocketsphinx/src/libpocketsphinx/fsg_lextree.c
xyzpocketsphinx/src/libpocketsphinx/fsg_search.c
xyzpocketsphinx/src/libpocketsphinx/hmm.c
xyzpocketsphinx/src/libpocketsphinx/kws_detections.c
xyzpocketsphinx/src/libpocketsphinx/kws_search.c
xyzpocketsphinx/src/libpocketsphinx/mdef.c
xyzpocketsphinx/src/libpocketsphinx/ms_gauden.c
xyzpocketsphinx/src/libpocketsphinx/ms_mgau.c
xyzpocketsphinx/src/libpocketsphinx/ms_senone.c
xyzpocketsphinx/src/libpocketsphinx/ngram_search.c
xyzpocketsphinx/src/libpocketsphinx/ngram_search_fwdflat.c
xyzpocketsphinx/src/libpocketsphinx/ngram_search_fwdtree.c
xyzpocketsphinx/src/libpocketsphinx/phone_loop_search.c
xyzpocketsphinx/src/libpocketsphinx/pocketsphinx.c
xyzpocketsphinx/src/libpocketsphinx/ps_alignment.c
xyzpocketsphinx/src/libpocketsphinx/ps_lattice.c
xyzpocketsphinx/src/libpocketsphinx/ps_mllr.c
xyzpocketsphinx/src/libpocketsphinx/ptm_mgau.c
xyzpocketsphinx/src/libpocketsphinx/s2_semi_mgau.c
xyzpocketsphinx/src/libpocketsphinx/state_align_search.c
xyzpocketsphinx/src/libpocketsphinx/tmat.c
xyzpocketsphinx/src/libpocketsphinx/vector.c
)
