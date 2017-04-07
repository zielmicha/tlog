#include "redis_conn.h"
#include "core/sleep.hh"

#include <sstream>

static const auto endline = std::string("\r\n");

static const auto set_prefix = std::string("*3") + endline + std::string("$3") + endline +
					std::string("set") + endline;

/** number of seconds to sleep before we try to reconnect to redis */
static const int REDIS_CONN_SLEEP_SEC = 4;

static const int RECONECT_NUM = 2;

/**
 * format a redis SET command
 */
std::string format_set(const char *key, int key_len, const char *val, int val_len) {
	std::stringstream ss;
	ss << set_prefix;

	// key
	ss << "$" << key_len  << endline;
	ss << std::string(key, key_len) + endline;

	// val
	ss << "$" << val_len << endline;
	ss << std::string(val, val_len) << endline;

	return ss.str();
}

redis_conn::redis_conn(socket_address sa, int idx)
	: _sa(sa)
	, _idx(idx)
{
	reconnect(RECONECT_NUM);
}

future<> redis_conn::reconnect(int retry_quota) {
		return connect(_sa).then([this, retry_quota] (connected_socket s) {
			_fd = std::move(s);
			_fd.set_keepalive(true);
			_out = _fd.output();
			_in = _fd.input();
			return make_ready_future<>();
		}).then_wrapped([this, retry_quota] (auto &&f) -> future<> {
			try {
				f.get();
			} catch(...) {
				std::cerr << "[ERROR]core "<< engine().cpu_id() << " failed to connect to redis " << _idx;
				std::cerr << ".retry quota = " << retry_quota << "\n";
				if (retry_quota > 0) {
					return sleep(std::chrono::seconds(REDIS_CONN_SLEEP_SEC)).then([this, retry_quota] {
							return this->reconnect(retry_quota - 1);
						});
				} 
				else {
					std::cerr << "[ERROR] tlog exiting because failed to connect to redis " << _idx << "\n";
					exit(1);
				}
			}
			return make_ready_future<>();
		});
}

future<bool> redis_conn::set(const std::string& key, const std::string& val) {
	return set(key.c_str(), key.length(), val.c_str(), val.length());
}

future<bool> redis_conn::set(const char *key, int key_len, 
		const char *val, int val_len) {

	auto data = format_set(key, key_len, val, val_len);
	return _out.write(data).then([this, key, key_len, val, val_len] {
			return _out.flush();
		}).then([this, key, key_len, val, val_len] {
			return _in.read();
		}).then([this, key, key_len, val, val_len] (auto buf) {
			auto str = std::string(buf.get(), buf.size());
			return make_ready_future<bool>(str=="+OK\r\n");
		}).then_wrapped([this, key, key_len, val, val_len] (auto &&f) -> future<bool>{
			auto catched = false;
			std::tuple<bool> ok;
			try {
				ok = f.get();
			} catch (...) {
				catched = true;
			}
			if (catched || !(std::get<0>(ok))) {
				std::cerr << "[ERROR]redis set failed.core "<< engine().cpu_id() << ".idx=" << _idx << "\n";
				return this->reconnect(RECONECT_NUM).then([this, key, key_len, val, val_len] {
						return this->set(key, key_len, val, val_len);
					});
			} else {
				return make_ready_future<bool>(true);
			}
		});
}

std::string format_get(const std::string& key) {
	std::stringstream ss;
	// prefix
	ss << "*2" << endline;

	// set command
	ss << "$3" << endline << "get" << endline;

	// key
	ss << "$" << key.length()  << endline;
	ss << key + endline;

	return ss.str();
}

future<bool> redis_conn::get(const std::string& key, uint8_t *val, unsigned int val_len) {
	auto data = format_get(key);
	auto p = net::packet(data.c_str(), data.length());
	return _out.write(std::move(p)).then([this, val, val_len] {
				return _out.flush();
			}).then([this, val, val_len] {
				return _in.read();
			}).then([this, val, val_len] (auto buf) {
				auto ok = false;
				auto prefix_len = 1 /* $ */ + floor(log10(val_len)) + 1 /* val_len */ + endline.length();
				if (buf && buf.size() == val_len + prefix_len) {
					std::memcpy(val, buf.get(), buf.size());
					ok = true;
				}
				return make_ready_future<bool>(ok);
			});
}

