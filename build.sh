#!/bin/bash
MAIN_PATH=`pwd`/src/main/
cd $MAIN_PATH
go get github.com/golang/glog
go install
