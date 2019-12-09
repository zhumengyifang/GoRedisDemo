package internal

import (
	"github.com/garyburd/redigo/redis"
)

/**
创建redis连接
*/
func CreateConn(protocolType string, redisAddress string, password string) (redis.Conn, error) {
	var conn redis.Conn
	var err error
	if password == "" {
		conn, err = redis.Dial(protocolType, redisAddress)
	} else {
		conn, err = redis.Dial(protocolType, redisAddress, redis.DialPassword(password))
	}
	if err != nil {
		return nil, err
	}
	return conn, nil
}
