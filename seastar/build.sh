g++-5 -std=c++14 `pkg-config --cflags --libs $SEASTAR/build/release/seastar.pc` -lpthread -lcapnp -lkj -lcapnp-rpc -lhiredis -lb2 -lisal main.cc tlog_schema.capnp.c++ -o main
