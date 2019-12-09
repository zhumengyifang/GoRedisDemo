package internal

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func GetAndSetDemo(conn redis.Conn) {
	_, err := conn.Do("Set", "demoKey", "demoValue")
	if err != nil {
		fmt.Println(err)
		return
	}

	value, err := redis.String(conn.Do("Get", "demoKey"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)
}

func MGetDemo(conn redis.Conn) {
	conn.Do("Set", "Key1", "Value1")
	conn.Do("Set", "Key2", "Value2")

	values, err := redis.Values(conn.Do("MGET", "Key1", "Key2"))
	if err != nil {
		fmt.Println(err)
		return
	}

	var value1 string
	var value2 string

	if _, err := redis.Scan(values, &value1, &value2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value1, value2)
}

func GetSetDemo(conn redis.Conn) {
	_, err := conn.Do("SET", "GetSetDemoK", "GetSetDemoV")
	if err != nil {
		fmt.Println(err)
		return
	}

	value, err := redis.String(conn.Do("GETSET", "GetSetDemoK", "replace"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)
}

func IncrDemo(conn redis.Conn) {
	if _, err := conn.Do("SET", "IncrKey", 1); err != nil {
		fmt.Println(err)
		return
	}

	if v, err := conn.Do("INCR", "IncrKey"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(v)
	}
}

func IncrByDemo(conn redis.Conn) {
	if _, err := conn.Do("SET", "IncrKey", 1); err != nil {
		fmt.Println(err)
		return
	}

	if v, err := conn.Do("INCRBY", "IncrKey", 5); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(v)
	}
}

func DecrbyDemo(conn redis.Conn) {
	if _, err := conn.Do("SET", "DecrbyKey", 1); err != nil {
		fmt.Println(err)
		return
	}

	if v, err := conn.Do("DECRBY", "DecrbyKey", 5); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(v)
	}
}

func AppendDemo(conn redis.Conn) {
	if _, err := conn.Do("SET", "AppendKey", "Go"); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := conn.Do("APPEND", "AppendKey", "Lang"); err != nil {
		fmt.Println(err)
		return
	}
}

func DelDemo(conn redis.Conn) {
	if _, err := conn.Do("SET", "DelKey", "DelValue"); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := conn.Do("DEL", "DelKey"); err != nil {
		fmt.Println(err)
	}
}

func ExDemo(conn redis.Conn) {
	if _, err := conn.Do("SET", "ExKey", "ExValue", "Ex", 10); err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(time.Second * 10)

	value, err := redis.String(conn.Do("Get", "ExKey"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)
}
