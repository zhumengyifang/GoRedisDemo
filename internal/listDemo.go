package internal

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func LPush(conn redis.Conn) {
	if _, err := conn.Do("LPUSH", "listDemo", "world"); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := conn.Do("LPUSH", "listDemo", "hello"); err != nil {
		fmt.Println(err)
		return
	}
	LRange(conn)
}

func LRange(conn redis.Conn) {
	if values, err := redis.Values(conn.Do("LRANGE", "listDemo", "0", "-1")); err != nil {
		fmt.Println(err)
		return
	} else {
		for _, v := range values {
			value := v.([]uint8)
			fmt.Println(B2S(value))
		}
	}
}

func LIndex(conn redis.Conn) {
	LPush(conn)
	fmt.Println("--------------------分割线--------------------")
	if value, err := redis.String(conn.Do("LINDEX", "listDemo", "0")); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(value)
	}

	if value, err := redis.String(conn.Do("LINDEX", "listDemo", "-1")); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(value)
	}

	if value, err := redis.String(conn.Do("LINDEX", "listDemo", "3")); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(value)
	}
}

func LInsert(conn redis.Conn) {
	if _, err := conn.Do("LPUSH", "listDemo", "world"); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := conn.Do("LPUSH", "listDemo", "hello"); err != nil {
		fmt.Println(err)
		return
	}

	if value, err := conn.Do("LINSERT", "listDemo", "BEFORE", "world", "---"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(value)
	}

	LRange(conn)
}

func RPush(conn redis.Conn) {
	if _, err := conn.Do("RPUSH", "listDemo", "hello"); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := conn.Do("RPUSH", "listDemo", "world"); err != nil {
		fmt.Println(err)
		return
	}
	LRange(conn)
}

func Rpoplpush(conn redis.Conn) {
	if _, err := conn.Do("RPUSH", "listDemo", "hello"); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := conn.Do("RPUSH", "listDemo", "foo"); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := conn.Do("RPUSH", "listDemo", "bar"); err != nil {
		fmt.Println(err)
		return
	}

	if value, err := redis.String(conn.Do("RPOPLPUSH", "listDemo", "myotherlist")); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(value)
	}
	LRange(conn)
}
