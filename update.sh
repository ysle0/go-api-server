#!/bin/bash

INIT_PATH="."
BUILD_NAME="go-api-server"
BRANCH="main"

PID=$(sudo lsof -i | grep $BUILD_NAME | awk '{print $2}')

echo $PID
kill -9 $PID

git checkout $BRANCH
git pull origin $BRANCH

cd $INIT_PATH
go build -o $BUILD_NAME

echo $PWD
nohup ./$BUILD_NAME &