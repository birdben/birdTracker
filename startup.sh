#!/bin/bash
LOG_PATH=`pwd`/logs
echo $LOG_PATH
if [ ! -d "$LOG_PATH" ]; then
    mkdir -p "$LOG_PATH"
fi
./bin/main -log_dir=$LOG_PATH
