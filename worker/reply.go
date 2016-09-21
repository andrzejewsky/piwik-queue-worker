package worker

import (
	"net/http"

	"github.com/andrzejewsky/tracker/server"
)

// HTTPReply replay http request readed from queue
type HTTPReply struct {
	host   string
	client *http.Client
}

// NewHTTPReply create a HTTPReply
func NewHTTPReply(host string) *HTTPReply {
	return &HTTPReply{
		host,
		&http.Client{},
	}
}

// Reply send fetched requests
func (h *HTTPReply) Reply(requestBus chan server.Request) {
	for request := range requestBus {
		h.send(request)
	}
}

func (h *HTTPReply) send(request server.Request) {
	req, _ := http.NewRequest(request.Method, h.host+request.URL.RequestURI(), request.Body)
	h.client.Do(req)
}
