@0xf4533cbae6e08506;

struct TlogResponse {
	status @0 :Int8;
	sequences @1 :List(UInt64);
	# only exist in flush response
}
struct TlogBlock {
	volumeId @0 :Text;
	sequence @1 :UInt64;
	lba @2 :UInt64;
	size @3 :UInt32;
	crc32  @4 :UInt32;
	data @5 :Data;
	timestamp @6 :UInt64;
}

struct TlogAggregation {
	name @0 :Text; # unused now
	size @1 :UInt64; # number of blocks in this aggregation
	timestamp @2 :UInt64;
	volumeId @3 :Text;
	blocks @4 :List(TlogBlock);
	prev @5 :Data; # hash of the previous aggregation
}
