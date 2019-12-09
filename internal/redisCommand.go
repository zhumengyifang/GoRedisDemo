package internal

const (

	RedisKeys  = "KEYS *"
	RedisDel   = "DEL"
	RedisClear = "FLUSHALL"

	//Hash
	RedisHashSet    = "HSET"
	RedisHashGet    = "HGET"
	RedisHashMSet   = "HMSET"
	RedisHashMGet   = "HMGET"
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
