#include "packet_cache.h"

void packet_cache::add(tlog_block *tb) {
	_packets.push(tb);
}
bool packet_cache::pick(std::queue<tlog_block *> *q, int flush_size) {
	if (this->size() == 0 || this->size() < flush_size) {
		return false;
	}
	
	for (int i=0; i  < flush_size && !_packets.empty(); i++) {
		auto tb = _packets.front();
		_packets.pop();
		q->push(tb);
	}
	return true;
}
