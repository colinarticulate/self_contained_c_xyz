add_library(libsphinxad
ad_alsa.c
ad_base.c
ad_openal.c
ad_oss.c
ad_pulse.c
ad_win32.c
)

target_include_directories(libsphinxad
PRIVATE ${CMAKE_CURRENT_SOURCE_DIR}
PUBLIC $(PROJECT_SOURCE_DIR)/include
)

target_link_libraries(libsphinxad sphinxbase_compiler_flags)
