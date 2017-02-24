# Transaction Log

## Settings
 - flush time: how maximum time we can wait data before force flush
 - max size: maximum database size before force flush

## Content
 - Name
 - Size
 - Timestamp
 - Blocks:
   - Size
   - Header
   - CRC
   - Data

## Workflow during flush
 - Hashing content (1) (If this is the first block, Hash (1) will be defined-private-key)
 - Compression
 - Encryption (Key (1))
 - Hashing encrypted block (2)
 - Erasure coding the full payload
 - Saving each parts to ForestDB ARDB Backend

