package internal

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type RedisLock struct {
	//保存连接
	Conn *redis.Conn
	//生成随机value用于判断是否为自己上锁
	uuid string
	//锁状态
	lock bool
}

/**
初始化
*/
func (redisLock *RedisLock) Init(protocolType string, ipAddress string, passWord string) error {
	conn, err := createConn(protocolType, ipAddress, passWord)
	if err != nil {
		return err
	}

	redisLock.Conn = &conn
	return nil
}

/**
创建redis连接
*/
func createConn(protocolType string, redisAddress string, password string) (redis.Conn, error) {
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

/**
尝试上锁
*/
func (redisLock *RedisLock) tryLock(key string) (bool, error) {
	//如果判断已经在锁定状态则直接返回
	if redisLock.lock {
		return true, nil
	}

	//生成uuid
	if redisLock.uuid == "" {
		redisLock.uuid = NewUUID()
	}

	//NX 若Key存在则无法修改其值
	//EX 过期时间 单位 秒
	result, err := (*redisLock.Conn).Do("SET", key, redisLock.uuid, "NX", "EX", 300)
	if err != nil {
		return false, err
	}

	if result == "OK" {
		redisLock.lock = true
		return redisLock.lock, nil
	}

	return false, nil
}

/**
释放锁
*/
func (redisLock *RedisLock) removeLock(key string) (bool, error) {
	if !redisLock.lock {
		return false, errors.New("Unlocked ")
	}

	//判断锁是否为自己使用
	script := "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end"

	//使用脚本方式
	result, err := redis.Bool((*redisLock.Conn).Do("EVAL", script, 1, key, redisLock.uuid))
	if err != nil {
		return result, err
	}

	return result, nil
}

/**
锁模版
*/
func (redisLock *RedisLock) LockForRedis(key string, f func()) error {
	result, err := redisLock.tryLock(key)
	if err != nil {
		return err
	}

	if !result {
		return errors.New("continue")
	}

	fmt.Println("tryLock")

	f()

	result, err = redisLock.removeLock(key)
	if err != nil {
		return err
	}

	if result {
		fmt.Println("removeLock")
	}

	return nil
}
