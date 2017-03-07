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

## TODO

- add options for objstor addr & port
- add options to specify K & M value of erasure encoding (currently hardcoded to 4 & 2)
- improve locking (if tlog server flush volume's tlog, it currently can't receive tlog 
  for that volume)
- add compression
- add encryption
- add proper logging
- fix segmentation fault issue (run it with `./main -c 1`, it stil has bug, but less)
