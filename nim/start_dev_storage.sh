#!/bin/bash
set -e

for i in {1..20}; do
    mkdir -p data/$i
    redis-server --port $((11000+i)) --dir $PWD/data/$i --pidfile $PWD/data/$i.pid --bind 127.0.0.1 --daemonize yes --appendonly yes
done
