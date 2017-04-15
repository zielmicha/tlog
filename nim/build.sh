#!/bin/sh
set -e
cd "$(dirname "$0")"

if [ -e nimenv.local ]; then
  echo 'nimenv.local exists. You may use `nimenv build` instead of this script.'
  #exit 1
fi

mkdir -p .nimenv/nim
mkdir -p .nimenv/deps

NIMHASH=9e199823be47cba55e62dd6982f02cf0aad732f369799fec42a4d8c2265c5167
if ! [ -e .nimenv/nimhash -a \( "$(cat .nimenv/nimhash)" = "$NIMHASH" \) ]; then
  echo "Downloading Nim http://nim-lang.org/download/nim-0.16.0.tar.xz (sha256: $NIMHASH)"
  wget http://nim-lang.org/download/nim-0.16.0.tar.xz -O .nimenv/nim.tar.xz
  if ! [ "$(sha256sum < .nimenv/nim.tar.xz)" = "$NIMHASH  -" ]; then
    echo "verification failed"
    exit 1
  fi
  echo "Unpacking Nim..."
  rm -r .nimenv/nim
  mkdir -p .nimenv/nim
  cd .nimenv/nim
  tar xJf ../nim.tar.xz
  mv nim-*/* .
  echo "Building Nim..."
  make -j$(getconf _NPROCESSORS_ONLN)
  cd ../..
  echo $NIMHASH > .nimenv/nimhash
fi

get_dep() {
  set -e
  cd .nimenv/deps
  name="$1"
  url="$2"
  hash="$3"
  srcpath="$4"
  new=0
  if ! [ -e "$name" ]; then
    git clone --recursive "$url" "$name"
    new=1
  fi
  if ! [ "$(cd "$name" && git rev-parse HEAD)" = "$hash" -a $new -eq 0 ]; then
     cd "$name"
     git fetch --all
     git checkout -q "$hash"
     git submodule update --init
     cd ..
  fi
  cd ../..
  echo "path: \".nimenv/deps/$name$srcpath\"" >> nim.cfg
}

echo "path: \".\"" > nim.cfg

get_dep capnp https://github.com/zielmicha/capnp.nim 69b9d07096a4b8af283a25e6a4bdf249e870351e ''
get_dep collections https://github.com/zielmicha/collections.nim 41a4c5451d2ad71bd28fbabb9335aed427b19ae4 ''
get_dep isa https://github.com/nimscale/isa 19a58e4e1be29b21300e0c7e324ba74322a075a8 ''
get_dep reactor https://github.com/zielmicha/reactor.nim ea5645fcfd693871e66d50bd163e77b9289058a9 ''

echo '# reactor.nim requires pthreads
threads: "on"

# enable debugging
passC: "-g"
passL: "-g"

verbosity: "0"
hint[ConvFromXtoItselfNotNeeded]: "off"
hint[XDeclaredButNotUsed]: "off"

debugger: "native"

threadanalysis: off # temporary, until asyncmacro is fixed

@if release:
  gcc.options.always = "-w -fno-strict-overflow"
  gcc.cpp.options.always = "-w -fno-strict-overflow"
  clang.options.always = "-w -fno-strict-overflow"
  clang.cpp.options.always = "-w -fno-strict-overflow"

  passC:"-ffunction-sections -fdata-sections -flto -fPIE"
  passL:"-Wl,--gc-sections -flto -fPIE"

  # tlog doesn'"'"'t handle input from untrusted sources (these checks are very cheap)
  obj_checks: off
  field_checks: off
  bound_checks: off
@end' >> nim.cfg

mkdir -p bin
ln -sf ../.nimenv/nim/bin/nim bin/nim

echo "building tlog"; bin/nim c -d:release --out:"$PWD/bin/tlog" tlog/main
