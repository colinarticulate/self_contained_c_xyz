#!/bin/bash

LIB_DIR="/home/dbarbera/Repositories/self_contained_c_xyz/Go_xyz_plus/build/ps_plus"
LIB_PATH="$LIB_DIR/libps_plus.so"
#CGO_LDFLAGS="-L${LIB_DIR} -lps_plus -Wl,-rpath,${LIB_DIR}" CGO_ENABLED=1 go build -buildmode=c-shared -o libcaller_plus.so main.go data.go strings.go
CGO_LDFLAGS="-L${LIB_DIR} -lps_plus -Wl,-rpath,${LIB_DIR}" CGO_ENABLED=1 go build -o caller_plus_exec main.go data.go strings.go

