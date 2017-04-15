#!/bin/bash
set -e
# -d:disableThreading
killall -INT main
nim c -d:profiling -d:release ./tlog/main
./tlog/main >/dev/null & DAEMONPID=$!
sleep 1
time ./tlog/bench_client 20000
killall -INT main
sleep 1
google-pprof --callgrind tlog/main profiler.out > callgrind.out.gperf
google-pprof --text tlog/main profiler.out
