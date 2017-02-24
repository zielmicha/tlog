# GIG TLOG 
Transaction log server

# [Specs](https://github.com/g8os/tlog/blob/master/SPECS.md)

## TLOG Syncronous Client

TLOG Client library is a client lib that works on GIG BLOCKSTOR (NBD Server)
Transaction log flushing should be syncronous with flushing to primary storage. If one of two operations (flushing to primary storage or to tlog) failed we should not consider another one as valid. 
TLOG Client communicates with TLOG Server via dpdk stack.
TLOG Server should store received log entries to fast NVME storage in a queue.

## TLOG Server
