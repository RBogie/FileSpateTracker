#!/bin/bash

#clean buildfolders
rm -rf protogen

#Generate protobuf files...
mkdir protogen/
cd proto/
protoc --go_out=../protogen/ *.proto
cd ..

#Build package
go install

EXIT_STATUS=$?

if [ $EXIT_STATUS == 0 ]; then
  echo "Build succeeded"
else
  echo "Build failed"
fi

exit $EXIT_STATUS
