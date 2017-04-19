#!/bin/bash
set -e
# -d:disableThreading
killall -INT main || true
#nim c -d:release ./tlog/main
valgrind --tool=callgrind --cache-sim=yes --branch-sim=yes ./tlog/main >/dev/null & DAEMONPID=$!
sleep 3
time ./tlog/bench_client 2000
kill $DAEMONPID
