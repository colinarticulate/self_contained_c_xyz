add_library(libxyzsphinxfe
fe_interface.c
fe_noise.c
fe_prespch_buf.c
fe_sigproc.c
fe_warp.c
fe_warp_affine.c
fe_warp_inverse_linear.c
fe_warp_piecewise_linear.c
fixlog.c
yin.c
)

target_include_directories(libxyzsphinxfe
PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}
PUBLIC $(PROJECT_SOURCE_DIR)/include
)

target_link_libraries(libxyzsphinxfe sphinxbase_compiler_flags)
