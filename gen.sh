#!/bin/bash
#
# Prerequisites:
#   protoc: https://github.com/golang/protobuf
#
# Usage:
#   ./gen.sh

mkdir -p /tmp/${USER}/tf
cd /tmp/${USER}/tf

# As of July 2018, the default supported version for Go is r1.7
# Check the output version of `hello_tf.go` in https://www.tensorflow.org/install/install_go
git clone -b r1.7 https://github.com/tensorflow/serving.git
git clone -b r1.7 https://github.com/tensorflow/tensorflow.git

mkdir -p generated_src

PROTOC_OPTS="-I tensorflow -I serving --go_out=plugins=grpc:generated_src"

protoc $PROTOC_OPTS serving/tensorflow_serving/apis/*.proto
protoc $PROTOC_OPTS serving/tensorflow_serving/config/*.proto
protoc $PROTOC_OPTS serving/tensorflow_serving/util/*.proto
protoc $PROTOC_OPTS serving/tensorflow_serving/sources/storage_path/*.proto
protoc $PROTOC_OPTS tensorflow/tensorflow/core/framework/*.proto
protoc $PROTOC_OPTS tensorflow/tensorflow/core/example/*.proto
protoc $PROTOC_OPTS tensorflow/tensorflow/core/lib/core/*.proto
protoc $PROTOC_OPTS tensorflow/tensorflow/core/protobuf/{saver,meta_graph}.proto

cd -

mv /tmp/${USER}/tf/generated_src src
