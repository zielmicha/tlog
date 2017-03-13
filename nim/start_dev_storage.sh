#!/bin/bash
set -e

for i in {1..10}; do
    mkdir -p data/$i
    redis-server --daemonize no --port $((11000+i)) --dir $PWD/data/$i --pidfile $PWD/data/$i.pid --bind 127.0.0.1 --daemonize yes
done
