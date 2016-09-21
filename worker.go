package main

import (
	"flag"

	"github.com/andrzejewsky/piwik-queue-worker/worker"
	"github.com/andrzejewsky/tracker/server"
)

var replyHost string
var redisAddr string
var redisPort string

func init() {
	flag.StringVar(&replyHost, "reply-host", "", "eg. http://localhost/")
	flag.StringVar(&redisAddr, "redis-addr", "", "redis host")
	flag.StringVar(&redisPort, "redis-port", "", "redis port")
}

func main() {
	flag.Parse()

	requestBus := make(chan server.Request)
	sem := make(chan bool)

	reply := worker.NewHTTPReply(replyHost)
	reader := worker.NewQueueFetcher(map[string]string{
		"host": redisAddr,
		"port": redisPort,
	})

	go reply.Reply(requestBus)
	go reader.StartFetching(requestBus)

	<-sem
}
