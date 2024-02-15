package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"projects/LDmitryLD/hugoproxy/proxy/config"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/service"
	pb "projects/LDmitryLD/hugoproxy/proxy/protos/gen/geogrpc"
	"projects/LDmitryLD/hugoproxy/proxy/rpc/geo"

	"google.golang.org/grpc"
)

const (
	grpcProtocol    = "grpc"
	rpcProtocol     = "rpc"
	jsonrpcProtocol = "json-rpc"
)

// func GetServerRPC(conf config.RPCServer, server *rpc.Server, geo *geo.GeoServiceGRPC) (Server, error) {

// 	switch conf.Type {
// 	case grpcProtocol:
// 		return NewServerGRPC(conf, geo), nil
// 	case rpcProtocol:
// 		return NewServerRPC(conf, server)
// 	case jsonrpcProtocol:
// 		return NewServerRPC(conf, server)
// 	default:
// 		return nil, fmt.Errorf("invalid protocol")
// 	}
// }

func GetServerRPC(conf config.RPCServer, geoService service.Georer) (Server, error) {

	switch conf.Type {
	case grpcProtocol:
		return NewServerGRPC(conf, geo.NewGeoServiceGRPC(geoService)), nil
	case rpcProtocol:
		geoRPC := geo.NewGeoServiceRPC(geoService)
		RPCServer := rpc.NewServer()
		err := RPCServer.Register(geoRPC)
		if err != nil {
			return nil, err
		}
		return NewServerRPC(conf, RPCServer)
	case jsonrpcProtocol:
		geoRPC := geo.NewGeoServiceRPC(geoService)
		RPCServer := rpc.NewServer()
		err := RPCServer.Register(geoRPC)
		if err != nil {
			return nil, err
		}
		return NewServerRPC(conf, RPCServer)
	default:
		return nil, fmt.Errorf("invalid protocol")
	}
}

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

type ServerGRPC struct {
	conf config.RPCServer
	srv  *grpc.Server
	geo  *geo.GeoServiceGRPC
}

func NewServerGRPC(conf config.RPCServer, geo *geo.GeoServiceGRPC) Server {
	return &ServerGRPC{
		conf: conf,
		srv:  grpc.NewServer(),
		geo:  geo,
	}
}

func (s *ServerGRPC) Serve(ctx context.Context) error {
	var err error

	chErr := make(chan error)
	go func() {
		var l net.Listener
		l, err = net.Listen("tcp", fmt.Sprintf(":%s", s.conf.Port))
		if err != nil {
			log.Println("gRPC server register error:", err)
			chErr <- err
		}

		log.Println("gRPC server started on ", s.conf.Port)

		pb.RegisterGeorerServer(s.srv, s.geo)

		if err = s.srv.Serve(l); err != nil {
			log.Println("grpc server error: ", err)
			chErr <- err
		}

	}()

	select {
	case <-chErr:
		return err
	case <-ctx.Done():
		s.srv.GracefulStop()
	}
	return err
}
