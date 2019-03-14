#!/bin/sh
./cmd-server/cmd-server -stderrthreshold=INFO -dbdb=m0v -dbuser=sridhar -dbpw=rcsp8 &
./cmd-gateway/cmd-gateway -stderrthreshold=INFO &


