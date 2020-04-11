package main

import (
	"fmt"
	"goredis/internal"
	"sync"
	"time"
)

const ipAddress = "192.168.170.137:6379"
const protocolType = "tcp"
const passWord = "myredis123"

var n = 0

func main() {
	fmt.Println("Start......")

	count := 200

	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			testLock()
			wg.Done()
		}()
		time.Sleep(time.Millisecond )
	}

	wg.Wait()

	fmt.Println(n)
	fmt.Println("End......")
}

func testLock() {
	key := "myLockKey"

	redisLock := new(internal.RedisLock)
	redisLock.Init(protocolType, ipAddress, passWord)

	result := redisLock.LockForRedis(key, test)
	if result != nil {
		fmt.Println(result)
	}
}

func test() {
	n = n + 1
}
