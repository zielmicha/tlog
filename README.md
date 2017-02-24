# GIG TLOG 
Transaction log server

# [Specs from Maxime](https://github.com/g8os/tlog/blob/master/SPECS.md)

## TLOG Syncronous Client

- TLOG Client library is a client lib that works on GIG BLOCKSTOR (NBD Server)
- Transaction log flushing should be syncronous with flushing to primary storage. If one of two operations (flushing to primary storage or to tlog) failed we should not consider another one as valid. 
- TLOG Client communicates with TLOG Server via dpdk stack.

## TLOG Server

- TLOG Server should store received log entries to fast NVME storage in a queue. 
- After storing log entry on NVME it replies to the client on successfull transaction.
- Separate worker stores log entries to OBJSTOR
- Assumption: We do not keep actual data, we keep only metadata log (if yes, then we do not need to keep [references in a storage](https://github.com/g8os/stor_client_lib/blob/master/README.md), remove it from there)
- We instantiate 10 instancecs of slow HDD using [OBJSTOR API](https://github.com/g8os/objstor)
- Each instance is used to keep a part of erasure coded data

We receive following information from the TLOG client:
- driveId
- lbaIndex
- hash2 (key of the data in data storage cluster)
- timestamp

Workflow 
