@0xf4533cbae6e08506;

struct TlogBlock {
	volumeId @0: UInt32;
	sequence @1: UInt32;
	lba @2: UInt64;
	size @3 :UInt32;
	crc32  @4 :UInt32;
	data @5 :Data;
	timestamp @6 :UInt64;
}

struct TlogAggregation {
	name @0 :Text;
	size @1: UInt64;
	timestamp @2 :UInt64;
	volumeId @3: UInt32;
	blocks @4 :List(TlogBlock);
}
