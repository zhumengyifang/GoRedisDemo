package internal

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func HSetAdnHGet(conn redis.Conn) {
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
	}else{
		var value1 string
		var value2 string

		if _, err := redis.Scan(values, &value1, &value2); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(value1, value2)
	}
}

