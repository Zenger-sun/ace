package ace

import (
	"fmt"
	"net/http"

	"github.com/lucas-clemente/quic-go/http3"
)

type Server struct {
	addr     string
	certPath string
	privPath string
	router   *Router
}

func (srv *Server) Serve() {
	if srv.addr == "" {
		panic(fmt.Errorf("server addr is nil"))
	}

	err := http3.ListenAndServeQUIC(srv.addr, srv.certPath, srv.privPath, nil)
	if err != nil {
		panic(err)
	}
}

func (srv *Server) AddRouter(router Router, handlerFunc HandlerFunc) {
	http.HandleFunc(router.Path, func(w http.ResponseWriter, r *http.Request) {
		err := handlerFunc(&TransPort{HttpWriter:w,HttpReq:r})
		if err != nil {
			router.Err("handler %v err: %v", router.Path, err)
		}
	})
}

func (srv *Server) Close() error {
	return nil
}

func (srv *Server) Network() string {
	return "http"
}

func (srv *Server) String() string {
	return srv.addr
}

func NewHttpServer(addr string, certPath, privPath string, router *Router) *Server {
	return &Server{
		addr:     addr,
		certPath: certPath,
		privPath: privPath,
		router:   router,
	}
}
