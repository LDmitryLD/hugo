package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"projects/LDmitryLD/hugoproxy/proxy/config"
)

const (
	rpcProtocol     = "rpc"
	jsonrpcProtocol = "json-rpc"
)

type ServerRPC struct {
	conf config.RPCServer
	srv  *rpc.Server
}

func NewServerRPC(conf config.RPCServer, server *rpc.Server) (Server, error) {
	switch conf.Type {
	case rpcProtocol:
		return &ServerRPC{conf: conf, srv: server}, nil
	case jsonrpcProtocol:
		return &ServerJSONRPC{conf: conf, srv: server}, nil
	default:
		return nil, fmt.Errorf("invalid protocol")
	}
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

type ServerJSONRPC struct {
	conf config.RPCServer
	srv  *rpc.Server
}

func (s *ServerJSONRPC) Serve(ctx context.Context) error {
	var err error

	chErr := make(chan error)
	go func() {
		var l net.Listener
		l, err = net.Listen("tcp", fmt.Sprintf(":%s", s.conf.Port))
		if err != nil {
			log.Println("json rpc server register error:", err)
			chErr <- err
		}

		log.Println("json rpc server started on ", s.conf.Port)
		var conn net.Conn
		for {
			select {
			case <-ctx.Done():
				log.Println("json rpc: stopping server")
				return
			default:
				conn, err = l.Accept()
				if err != nil {
					log.Println("json rpc: net tcp accept error:", err)
				}
				go s.srv.ServeCodec(jsonrpc.NewServerCodec(conn))
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
