#!/bin/bash

BASEDIR=${BASEDIR:-`pwd`}
ROOT_DIR=${ROOT_DIR:-`dirname $BASEDIR`}
SRC_DIR=${SRC_DIR:-"$ROOT_DIR/src/internal"}

DOCKER=${DOCKER:-`which docker`}

cd $SRC_DIR
echo "#################################################"
echo "#  BASEDIR: $BASEDIR"
echo "#  ROOT_DIR: $ROOT_DIR"
echo "#  SRC_DIR: $SRC_DIR"
echo "#"
echo "#  CURRENT DIR: $PWD"
echo "#################################################"

$DOCKER build -t go-rest-api -f $BASEDIR/Dockerfile .