#!/bin/bash
set -e
# -d:disableThreading
[ "$SEASTAR" = "" ] && nim c -d:release tlog/main.nim

bench() {
    if [ "$SEASTAR" = "" ]; then
        ./tlog/main & DAEMONPID=$!
        # about 0.9 s on my laptop
    else
        ../seastar/main --objstor_port 11001  & DAEMONPID=$!
        sleep 1
    fi
    trap "kill -9 $DAEMONPID" exit
    sleep 1
    time ./tlog/bench_client 20000
    killall main
}

for i in $(seq 6); do
    bench
done
