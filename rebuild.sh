set -ex
mkdir -p build
cd build
rm -rf *
cmake .. -DTAF_WEB_HOST=http://taf.test.whup.com:8080
make
make tar