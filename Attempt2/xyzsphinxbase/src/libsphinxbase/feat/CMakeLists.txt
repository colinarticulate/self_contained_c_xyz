add_library(libxyzsphinxfeat
agc.c
cmn.c
cmn_live.c
feat.c
lda.c
)

target_include_directories(libxyzsphinxfeat
PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}
PUBLIC $(PROJECT_SOURCE_DIR)/include
)

target_link_libraries(libxyzsphinxfeat sphinxbase_compiler_flags)