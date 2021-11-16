package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// Server is a gateway server.
type Server struct {
	*fasthttp.Server

	log  *log.Helper
	addr string
}

// New new a gateway server.
func New(logger log.Logger, handler http.Handler, addr string, timeout time.Duration, idleTimeout time.Duration) *Server {

	srv := &Server{
		addr: addr,
		Server: &fasthttp.Server{
			// Addr: addr,
			Handler:     fasthttpadaptor.NewFastHTTPHandler(handler),
			ReadTimeout: timeout,
			// ReadHeaderTimeout: timeout,
			WriteTimeout: timeout,
			IdleTimeout:  idleTimeout,
		},
		log: log.NewHelper(logger),
	}
	return srv
}

// Start start the server.
func (s *Server) Start(ctx context.Context) error {
	s.log.Infof("server listening on %s", s.addr)
	// s.BaseContext = func(net.Listener) context.Context {
	// 	return ctx
	// }

	return s.ListenAndServe(s.addr)
}

// Stop stop the server.
func (s *Server) Stop(ctx context.Context) error {
	s.log.Info("server stopping")
	return s.Shutdown()
}
