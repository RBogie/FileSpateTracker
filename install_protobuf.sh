#!/bin/sh

set -ex
wget https://github.com/google/protobuf/archive/v3.0.0-beta-1.tar.gz
tar -xzvf v3.0.0-beta-1.tar.gz
cd v3.0.0-beta-1 && ./configure --prefix=/usr && make && sudo make install
