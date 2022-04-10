package main

import (
	"fmt"

	redis "github.com/go-redis/redis"
)

// Monitor monitor struct, to monitor redis
type Monitor struct {
	client *redis.Client
}

// Dispose disose monitor instance
func (monitor *Monitor) Dispose() {
	monitor.client.Close()
}

// NewMonitor create monitor instance
func NewMonitor(client *redis.Client) *Monitor {
	monitor := new(Monitor)
	monitor.client = client
	return monitor
}

// Memory check redis memory used
func (monitor *Monitor) Memory(key string) (int64, error) {
	return monitor.client.MemoryUsage(key).Result()
}

// Info get redis server basic information
func (monitor *Monitor) Info() (string, error) {
	return monitor.client.Info([]string{}...).Result()
}

// Refcount get the given key references count
func (monitor *Monitor) Refcount(key string) (int64, error) {
	return monitor.client.ObjectRefCount(key).Result()
}

// Events registry a event into redis
func (monitor *Monitor) Events() {
	monitor.client.ConfigSet("notify-keyspace-events", "Ex")
}

// 功能使用
func monitor() {
	monitor := NewMonitor(NewRedisClient())
	defer monitor.Dispose()

	// info
	fmt.Println("redis info:")
	fmt.Println(monitor.Info())

	// object refcount
	fmt.Println("object ref count:")
	fmt.Println(monitor.Refcount("hello"))
}
