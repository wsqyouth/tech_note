package main

import (
	"fmt"
	"time"

	redis "github.com/go-redis/redis"
)

// StreamPubSub stream struct
type StreamPubSub struct {
	client *redis.Client
}

// NewStreamPubSub create a stream instance
func NewStreamPubSub(client *redis.Client) *StreamPubSub {
	return &StreamPubSub{client: client}
}

// Publish publish a redis stream message
func (streamps *StreamPubSub) Publish(stream string, id string, values map[string]interface{}) {
	arg := redis.XAddArgs{Stream: stream, ID: id, Values: values}
	streamps.client.XAdd(&arg)
}

// Consumer consume redis stream message
func (streamps *StreamPubSub) Consumer(stream string, id string) ([]redis.XMessage, error) {
	result, err := streamps.client.XRead(&redis.XReadArgs{
		Streams: []string{stream, id},
		Count:   2,
		Block:   2 * time.Second,
	}).Result()

	if err != nil {
		return nil, err
	}

	return result[0].Messages, nil
}

//功能使用
// Stream test redis stream feature
func Stream() {
	client1 := NewRedisClient()
	defer client1.Close()
	client2 := NewRedisClient()
	defer client2.Close()
	streamPub := NewStreamPubSub(client1)
	streamSub := NewStreamPubSub(client2)

	wait := make(chan int)
	// consumer
	go func(chan int) {
		msg, err := streamSub.Consumer("stream", "1")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("stream message received:", msg)
		}
		wait <- 1
	}(wait)

	// publish msg
	go func() {
		// time.Sleep(3 * time.Second)
		streamPub.Publish("stream", "1-0", map[string]interface{}{"action": "moving"})
		streamPub.Publish("stream", "2-0", map[string]interface{}{"laugh": "hahaha"})
		streamPub.Publish("stream", "3-0", map[string]interface{}{"cry": "wuwuwu"})
		streamPub.Publish("stream", "4-0", map[string]interface{}{"talk": "gogogo"})
	}()

	<-wait

}
