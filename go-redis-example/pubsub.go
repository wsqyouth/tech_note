package main

import (
	"fmt"
	"time"

	redis "github.com/go-redis/redis"
)

// PubSub redis pubsub struct
type PubSub struct {
	client *redis.Client
}

// Dispose dispose pubsub instance
func (pubsub *PubSub) Dispose() {
	pubsub.client.Close()
}

// NewPubsub create pubsub instance
func NewPubsub(client *redis.Client) *PubSub {
	return &PubSub{client: client}
}

// Publish pubslish a message into the channel
func (pubsub *PubSub) Publish(channel string, msg interface{}) (int64, error) {
	return pubsub.client.Publish(channel, msg).Result()
}

// Subscribe subscribe a redis channel to receive message
func (pubsub *PubSub) Subscribe(channel string, msg chan string) {
	sub := pubsub.client.Subscribe(channel)
	go func() {
		for {
			outPut, _ := sub.ReceiveMessage()
			msg <- outPut.Payload
		}
	}()
}

// 功能使用

// Pubsub publish|subscribe
func Pubsub() {
	pub := NewPubsub(NewRedisClient())
	sub := NewPubsub(NewRedisClient())

	msgChan := make(chan string)
	go func() {
		sub.Subscribe("talk", msgChan)
		for {
			msg := <-msgChan
			fmt.Println(msg)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	pub.Publish("talk", "How are you?")
	pub.Publish("talk", "fine, thank you, and you?")
	pub.Publish("talk", "i am fine too!")
	pub.Publish("talk", "how old are you?")
}
