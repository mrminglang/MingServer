
execute_process(COMMAND go env GOPATH OUTPUT_VARIABLE GOPATH)
string(REGEX REPLACE "\n$" "" GOPATH "${GOPATH}")

include(${CMAKE_CURRENT_SOURCE_DIR}/cmake/taf-tools.cmake)

cmake_minimum_required(VERSION 2.8)

project(MingServer Go) # select GO compile

gen_server(MingApp MingServer)

add_subdirectory(client)

# go env -w GO111MODULE=on
# go mod init
# mkdir build
# cd build
# cmake ..
# make