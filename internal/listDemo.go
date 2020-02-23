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

/**
Blpop 命令移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
*/
func Blpop(conn redis.Conn) {
	defer conn.Close()
	LPush(conn)
	if values, err := redis.Values(conn.Do("BLPOP", "listDemo", 100)); err != nil {
		fmt.Println(err)
		return
	} else {
		for _, v := range values {
			value := v.([]uint8)
			fmt.Println(B2S(value))
		}
	}

	LRange(conn)
}

/**
Brpop 命令移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
*/
func Brpop(conn redis.Conn) {
	defer conn.Close()
	LPush(conn)
	if values, err := redis.Values(conn.Do("BRPOP", "listDemo", 100)); err != nil {
		fmt.Println(err)
		return
	} else {
		for _, v := range values {
			value := v.([]uint8)
			fmt.Println(B2S(value))
		}
	}

	LRange(conn)
}

/**
Brpoplpush 命令从列表中弹出一个值，将弹出的元素插入到另外一个列表中并返回它； 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
*/
func Brpoplpush(conn redis.Conn) {
	defer conn.Close()
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

	if value, err := redis.String(conn.Do("BRPOPLPUSH", "listDemo", "myotherlist", 100)); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(value)
	}
	LRange(conn)
}

func Lrem(conn redis.Conn) {
	defer conn.Close()
	RPush(conn)
	if _, err := conn.Do("LREM", "listDemo", 0, "world"); err != nil {
		fmt.Println(err)
		return
	}

	LRange(conn)
}

func Llen(conn redis.Conn) {
	defer conn.Close()
	RPush(conn)

	if value, err := conn.Do("LLEN", "listDemo"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(value)
	}
}

/**
Ltrim 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。

下标 0 表示列表的第一个元素，以 1 表示列表的第二个元素，以此类推。 你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。
*/
func Ltrim(conn redis.Conn) {
	defer conn.Close()
	RPush(conn)

	if _, err := conn.Do("LTRIM", "listDemo",1, -1); err != nil {
		fmt.Println(err)
		return
	}

	LRange(conn)
}
