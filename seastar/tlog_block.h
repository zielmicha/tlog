#ifndef __TLOG_BLOCK_H_
#define __TLOG_BLOCK_H_

struct tlog_block {
	uint32_t _vol_id;
	uint64_t _sequence;
	uint64_t _lba;
	uint32_t _size;
	uint32_t _crc;
	uint8_t *_data;
	uint64_t _timestamp;
public:
	tlog_block(uint32_t vol_id, uint64_t seq, uint64_t lba, uint32_t size, uint32_t crc,
			uint8_t *data, uint64_t timestamp) {
		_vol_id = vol_id;
		_sequence = seq;
		_lba = lba;
		_size = size;
		_crc = crc;
		_data = (uint8_t *) malloc (size);
		std::memcpy(_data, data, size);
		_timestamp = timestamp;
	}
	~tlog_block() {
		free(_data);
	}
};
#endif
