package worker

import (
	"encoding/json"

	"github.com/andrzejewsky/go-queue/queues"
	"github.com/andrzejewsky/tracker/server"
)

// QueueFetcher reading stored requests in the queue
type QueueFetcher struct {
	queue queues.Queue
}

// NewQueueFetcher new instance of NewQueueFetcher
func NewQueueFetcher(config map[string]string) *QueueFetcher {

	queue, _ := queues.GetQueue("redis", config)

	return &QueueFetcher{queue}
}

// StartFetching start reading requests from the queue
func (f *QueueFetcher) StartFetching(requestBus chan server.Request) {
	for {
		plainRequest, _ := f.queue.Pop()
		var request server.Request
		json.Unmarshal([]byte(plainRequest), &request)
		requestBus <- request
	}
}
