#!/bin/bash
set -ex
mkdir -p build
cd build
cmake ..
make
cd -
./build/bin/MingServer --config=config/config.conf
