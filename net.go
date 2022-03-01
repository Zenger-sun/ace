package ace

import (
	"net"
	"net/http"
)

const (
	NET_HTTP = iota
	NET_GRPC
	NET_STREAM
)

type Listener interface {
	net.Addr
	Serve()
	AddRouter(Router, HandlerFunc)
	Close() error
}

type HandlerFunc func(trans *TransPort) error

type Router struct {
	*Context
	Method string
	Path   string
	Handle HandlerFunc
}

type TransPort struct {
	HttpWriter http.ResponseWriter
	HttpReq    *http.Request
}
