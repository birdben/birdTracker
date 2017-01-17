#!/bin/bash
go test jsonlog -v -log_dir=`pwd`/logs -stderrthreshold=INFO
