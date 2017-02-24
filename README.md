# GIG TLOG 
Transaction log server

# [Specs from Maxime](https://github.com/g8os/tlog/blob/master/SPECS.md)

## TLOG Client

- TLOG Client library is a client lib that works on GIG BLOCKSTOR (NBD Server)
- TLOG Client communicates with TLOG Server via dpdk stack.

### Data send from the client to server:
- volume ID
- LBA
- data
- timestamp

## TLOG Server

- TLOG Server should store received log entries and aggregate them in a capnp structure in memmory
- After storing log entry it replies to the client on successfull transaction.
- after timeout or size of the aggregation is reached, we flush to OBJSTOR on top of ardb with forstdb datastore.
- Ideal setup would be to spread erasure coded blocks on different disks.

### Settings

flush time: how maximum time we can wait data before force flush  
max size: maximum database size before force flush

### Enrty log structure:
```
Name
Size
Timestamp
Blocks:
  Size
  Header
  CRC
  Data
```

## Workflow during flush

1. Hashing content (1) (blockhain style)
 - If this is the first block, Hash (1) will be defined-private-key, 
   othwise it's the has of the previous log entry.
2. Compression
3. Encryption (Key (1))
4. Hashing encrypted block (2)
5. Erasure coding the full payload
6. Saving each parts to ForestDB ARDB Backend
