add_library(libxyzsphinxlm
fsg_model.c
jsgf.c
jsgf_parser.c
jsgf_parser.y
jsgf_scanner.c
lm_trie.c
lm_trie_quant.c
ngram_model.c
ngram_model_set.c
ngram_model_trie.c
ngrams_raw.c
)

target_include_directories(libxyzsphinxlm
PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}
PUBLIC $(PROJECT_SOURCE_DIR)/include
)

target_link_libraries(libxyzsphinxlm sphinxbase_compiler_flags)