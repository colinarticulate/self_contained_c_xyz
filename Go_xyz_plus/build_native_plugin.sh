#!/bin/bash

LIB_DIR="/home/dbarbera/Repositories/self_contained_c_xyz/Go_xyz_plus/build/ps_plus"
LIB_PATH="$LIB_DIR/libps_plus.so"
CGO_LDFLAGS="-L$LIB_DIR" LD_LIBRARY_PATH="$LIB_DIR" go build -o caller_plus/caller_plus_exec caller_plus/main.go caller_plus/data.go caller_plus/strings.go

