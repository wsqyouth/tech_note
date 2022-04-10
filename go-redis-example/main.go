package main

import (
	"fmt"
	"net"

	"github.com/wsqyouth/tech_note/go-redis-example/redisprotocol"
)

func main() {
	client := NewRedisClient()
	defer client.Close()

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("redis is reachable")
	// monitor
	// monitor()

	// listen expire
	//EventListen()

	// pubsub
	//Pubsub()

	// stream
	//Stream()

	// protocol
	RedisProtocol()

	// wait
	wait := make(chan int)
	<-wait
}

// RedisProtocol test redis protocol
func RedisProtocol() {
	conn, _ := net.Dial("tcp", "127.0.0.1:6379")
	srp := redisprotocol.NewSimpleReidsPotocol(conn)
	go func() {
		cmd := redisprotocol.NewCmd(redisprotocol.Cmd_Get, []interface{}{"person"})
		srp.Send(redisprotocol.ICmd(cmd))
	}()
	go func() {
		srp.Receive()
	}()
}
