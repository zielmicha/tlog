# Tlog Seastar Server

## Build Seastar framework

Build seastar with DPDK
```
cd
git clone https://github.com/scylladb/seastar.git
cd seastar
apt-get update
./install-dependencies.sh
git submodule update --init
./configure.py --compiler=g++-5 --enable-dpdk
ninja
```

Use `ninja -j X` if you ran out of memory.

export SEASTAR env var

```
export SEASTAR=YOUR_SEASTAR_DIR
```

## Build tlog server

```
./build.sh
```

## Run in development environment

This guide assuming user to use default options

### options

Tlog server options can be seen by executing `./main -h` and look at `App options` sections
```
$ ./main -h
App options:
  -h [ --help ]                         show help message
  --port arg (=11211)                   tcp port
  --flush_size arg (=25)                flush_size
  --flush_time arg (=25)                flush_time (seconds)
  --k arg (=4)                          K variable of erasure encoding
  --m arg (=2)                          M variable of erasure encoding
  --objstor_addr arg (=127.0.0.1)       objstor address
  --objstor_port arg (=16379)           objstor first port
  --priv_key arg (=my-secret-key)       private key

```

### start 1+K+M number of ardb server in localhost

first server for metadata, it store `last hash` value of each volume ID. Listen in port `16379`.

Other (6 == k+m == 4+2) servers listens from port 16380-16385 store the erasure encoded data

### Start tlog server 

```
./main
```
### Use provided client library to send tlog data

Client lib can be found in [client dir](https://github.com/g8os/tlog/tree/master/client).

Usage example can be found in [examples dir](https://github.com/g8os/tlog/tree/master/client/examples)

## TODO

- add crc
- add compression
- add encryption
- add proper logging
