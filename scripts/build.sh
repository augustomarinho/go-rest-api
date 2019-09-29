#!/bin/bash

BASEDIR=${BASEDIR:-`pwd`}
ROOT_DIR=${ROOT_DIR:-`dirname $BASEDIR`}
SRC_DIR=${SRC_DIR:-"$ROOT_DIR/src/internal/app"}
TEST_DIR=${TEST_DIR:-"$ROOT_DIR/src/internal/test"}
BIN_DIR=${BIN_DIR:-"$ROOT_DIR/bin"}

GO=${GO:-`which go`}

cd $SRC_DIR

echo "Building project"
$GO build -v
echo "Project Built"

echo "Running tests"
$GO test -v $TEST_DIR/...