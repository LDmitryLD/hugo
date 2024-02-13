package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"projects/LDmitryLD/hugoproxy/proxy/config"
)

type ServerRPC struct {
	conf config.RPCServer
	srv  *rpc.Server
}

func NewServerRPC(conf config.RPCServer, server *rpc.Server) Server {
	return &ServerRPC{conf: conf, srv: server}
}

func (s *ServerRPC) Serve(ctx context.Context) error {
	var err error

	chErr := make(chan error)
	go func() {
		var l net.Listener
		l, err = net.Listen("tcp", fmt.Sprintf(":%s", s.conf.Port))
		if err != nil {
			log.Println("rpc server register error")
			chErr <- err
		}

		log.Println("rpc server started on ", s.conf.Port, "  ", l.Addr())
		var conn net.Conn
		for {
			select {
			case <-ctx.Done():
				log.Println("rpc: stopping server")
				return
			default:

				conn, err = l.Accept()
				if err != nil {
					log.Println("json rpc: net tcp accpet error:", err)
				}
				go s.srv.ServeConn(conn)
			}
		}
	}()

	select {
	case <-chErr:
		return err
	case <-ctx.Done():
	}

	return err
}
