package main

import (
	"fmt"
	"time"
)

const (
	// REGISTRYEVENT registry event
	REGISTRYEVENT string = "notify-keyspace-events"
	// EVENTTYPE event type is expire
	EVENTTYPE = "Ex"
	// EXPIREDEVENT expired event name
	EXPIREDEVENT string = "__keyevent@0__:expired"
)

// EventHandler event hdnaler obejct
type EventHandler struct{}

// NewEventHandler create new event handler instance
func NewEventHandler() *EventHandler {
	return new(EventHandler)
}

// ListenEventExpired listen the expired event
func (*EventHandler) ListenEventExpired(expiredKey chan string) {
	// publish
	NewRedisClient().ConfigSet(REGISTRYEVENT, EVENTTYPE)
	// subscribe
	sub := NewRedisClient().Subscribe(EXPIREDEVENT)
	// listen events
	go func() {
		for {
			msg, _ := sub.ReceiveMessage()
			if msg != nil {
				expiredKey <- msg.Payload
			}
		}
	}()
}

// 功能使用

// EventListen test redis event feature via pub/sub
func EventListen() {
	eventHandler := NewEventHandler()
	expiredChan := make(chan string)
	eventHandler.ListenEventExpired(expiredChan)
	fmt.Println("listen expired event...")
	go func() {
		for {
			msg := <-expiredChan
			fmt.Println("expired key", msg)
			time.Sleep(100 * time.Millisecond)
		}
	}()
}
