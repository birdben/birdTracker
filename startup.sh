#!/bin/bash
LOG_PATH=`pwd`/logs
echo "日志目录："$LOG_PATH
if [ ! -d "$LOG_PATH" ]; then
    mkdir -p "$LOG_PATH"
fi
./bin/main -v 10 -log_dir=$LOG_PATH -stderrthreshold=INFO
