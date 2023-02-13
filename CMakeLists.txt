# c_xyz cmake file

cmake_minimum_required(VERSION 3.4.1)

project(PS_PLUS)

# set(SPHINXBASE_DIR /home/dbarbera/Repositories/sefcontained_c_xyz/xyzsphinxbase)
# set(POCKETSPHINX_DIR /home/dbarbera/Repositories/sefcontained_c_xyz/xyzpocketsphinx)

#Just need to add ALL the source files

add_executable(ps_plus ps_plus.cpp)
