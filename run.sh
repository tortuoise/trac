#!/bin/sh
./cmd-server/cmd-server -stderrthreshold=INFO &
./cmd-gateway/cmd-gateway -stderrthreshold=INFO &


