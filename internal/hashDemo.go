package internal

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func HSetAndHGet(conn redis.Conn) {
	if _, err := conn.Do("HSET", "hashDemo", "name", "xiaowang", "age", 20); err != nil {
		fmt.Println(err)
		return
	}

	if v, err := redis.String(conn.Do("HGET", "hashDemo", "name")); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(v)
	}

	if v, err := redis.String(conn.Do("HGET", "hashDemo", "age")); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(v)
	}
}

/**
Redis Hmset 命令用于同时将多个 field-value (字段-值)对设置到哈希表中。
此命令会覆盖哈希表中已存在的字段。
如果哈希表不存在，会创建一个空哈希表，并执行 HMSET 操作
*/
func HMSetAndHMGet(conn redis.Conn) {
	if _, err := conn.Do("HMSET", "hashDemo", "name", "xiaowang", "age", 20); err != nil {
		fmt.Println(err)
		return
	}

	if values, err := redis.Values(conn.Do("HMGET", "hashDemo", "name", "age")); err != nil {
		fmt.Println(err)
		return
	} else {
		var value1 string
		var value2 string

		if _, err := redis.Scan(values, &value1, &value2); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(value1, value2)
	}
}

func HGetAllI(conn redis.Conn) {
	if _, err := conn.Do("HMSET", "hashDemo", "name", "xiaowang", "age", 20); err != nil {
		fmt.Println(err)
		return
	}

	if values, err := redis.Values(conn.Do("HGETALL", "hashDemo")); err != nil {
		fmt.Println(err)
		return
	} else {
		for _, v := range values {
			value := v.([]uint8)
			fmt.Println(B2S(value))
		}
	}
}

func B2S(bs []uint8) string {
	var ba []byte
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}

func Hexists(conn redis.Conn) {
	if _, err := conn.Do("HMSET", "hashDemo", "name", "xiaowang", "age", 20); err != nil {
		fmt.Println(err)
		return
	}

	if value, err := redis.Bool(conn.Do("HEXISTS", "hashDemo", "name")); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(value)
	}

	if value, err := redis.Bool(conn.Do("HEXISTS", "hashDemo", "fullname")); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(value)
	}
}

func HLen(conn redis.Conn) {
	if _, err := conn.Do("HMSET", "hashDemo", "name", "xiaowang", "age", 20); err != nil {
		fmt.Println(err)
		return
	}

	if value, err := redis.Int(conn.Do("HLEN", "hashDemo")); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(value)
	}
}

func HKeys(conn redis.Conn) {
	if _, err := conn.Do("HMSET", "hashDemo", "name", "xiaowang", "age", 20); err != nil {
		fmt.Println(err)
		return
	}

	if values, err := redis.Values(conn.Do("HKEYS", "hashDemo")); err != nil {
		fmt.Println(err)
		return
	} else {
		for _, v := range values {
			value := v.([]uint8)
			fmt.Println(B2S(value))
		}
	}
}

func HValues(conn redis.Conn) {
	if _, err := conn.Do("HMSET", "hashDemo", "name", "xiaowang", "age", 20); err != nil {
		fmt.Println(err)
		return
	}

	if values, err := redis.Values(conn.Do("HVALS", "hashDemo")); err != nil {
		fmt.Println(err)
		return
	} else {
		for _, v := range values {
			value := v.([]uint8)
			fmt.Println(B2S(value))
		}
	}
}

func HDel(conn redis.Conn) {
	if _, err := conn.Do("HMSET", "hashDemo", "name", "xiaowang", "age", 20); err != nil {
		fmt.Println(err)
		return
	}

	if value, err := redis.Int(conn.Do("HDEL", "hashDemo", "name")); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(value)
	}

	if value, err := redis.Int(conn.Do("HDEL", "hashDemo", "age")); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(value)
	}
}
