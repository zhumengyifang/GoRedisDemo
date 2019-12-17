package internal

const (
	RedisClear = "FLUSHALL"

	RedisHashGetAll = "HGETALL"
	RedisHashEXISTS = "HEXISTS"
	RedisHashLen    = "HLEN"
	RedisHashKeys   = "HKEYS"
	RedisHashVals   = "HVALS"
	RedisHashDel    = "HDEL"

	//List
	RedisListPush  = "LPUSH"
	RedisListRange = "LRANGE"
	RedisListPop   = "LPOP"
	RedisListLen   = "LLEN"
)
