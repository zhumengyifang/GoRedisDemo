package main

import (
	"fmt"
	"goredis/internal"
)

const ipAddress = "192.168.1.8:6379"
const protocolType = "tcp"
const passWord = "myredis123"

func main() {
	fmt.Println("Start......")
	conn, err := internal.CreateConn(protocolType, ipAddress, passWord)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	internal.Rpoplpush(conn)

	fmt.Println("End......")
}
