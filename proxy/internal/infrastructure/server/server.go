package server

import (
	"context"
	"log"
	"net/http"
	"projects/LDmitryLD/hugoproxy/proxy/config"
	"time"
)

type Server interface {
	Serve(ctx context.Context) error
}

type HTTPServer struct {
	conf config.Server
	srv  *http.Server
}

func NewHTTPServer(conf config.Server, server *http.Server) Server {
	return &HTTPServer{conf: conf, srv: server}
}

func (s *HTTPServer) Serve(ctx context.Context) error {
	var err error

	chErr := make(chan error)

	go func() {
		log.Println("server started")
		if err = s.srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Println("http listen and server error:", err)
			chErr <- err
		}
	}()

	select {
	case <-chErr:
		return err
	case <-ctx.Done():
	}

	ctxShutdown, cancel := context.WithTimeout(context.Background(), s.conf.ShutdoundTimeout*time.Second)
	defer cancel()
	err = s.srv.Shutdown(ctxShutdown)

	return err
}
