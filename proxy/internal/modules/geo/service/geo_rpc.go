package service

import (
	"context"
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"projects/LDmitryLD/hugoproxy/proxy/config"
	"projects/LDmitryLD/hugoproxy/proxy/internal/models"
	pb "projects/LDmitryLD/hugoproxy/proxy/protos/gen/geogrpc"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcProtocol    = "grpc"
	rpcProtocol     = "rpc"
	jsonrpcProtocol = "json-rpc"
)

func GetlientRPC(protocol string, conf config.GeoRPC) (Georer, error) {
	switch protocol {
	case rpcProtocol:
		client, err := newClient(conf, protocol)
		if err != nil {
			return nil, err
		}
		return NewGeoRPCClient(client), nil
	case jsonrpcProtocol:
		client, err := newClient(conf, protocol)
		if err != nil {
			return nil, err
		}
		return NewGeoRPCClient(client), nil
	case grpcProtocol:
		conn, err := grpc.Dial(fmt.Sprintf("%s:%s", conf.Host, conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Println("grpc server connect error:", err)
			return nil, err
		}

		client := pb.NewGeorerClient(conn)
		return NewGeoGRPCCLient(client), nil
	default:
		return nil, fmt.Errorf("invalid protocol")
	}
}

type GeoRPCClient struct {
	client *rpc.Client
}

func NewGeoRPCClient(client *rpc.Client) *GeoRPCClient {
	return &GeoRPCClient{
		client: client,
	}
}

func (g *GeoRPCClient) SearchAddresses(in SearchAddressesIn) SearchAddressesOut {
	var out SearchAddressesOut
	err := g.client.Call("GeoServiceRPC.SearchAddresses", in, &out)
	if err != nil {
		out.Err = err
	}

	return out
}

func (g *GeoRPCClient) GeoCode(in GeoCodeIn) GeoCodeOut {
	var out GeoCodeOut
	err := g.client.Call("GeoServiceRPC.GeoCode", in, &out)
	if err != nil {
		out.Err = err
	}

	return out
}

type GeoGRPCClient struct {
	client pb.GeorerClient
}

func NewGeoGRPCCLient(client pb.GeorerClient) *GeoGRPCClient {
	return &GeoGRPCClient{
		client: client,
	}
}

func (g *GeoGRPCClient) SearchAddresses(in SearchAddressesIn) SearchAddressesOut {
	res, err := g.client.SearchAddresses(context.Background(), &pb.SearchAddressesRequest{Query: in.Query})

	address := models.Address{
		Lat: res.Address.Lat,
		Lon: res.Address.Lon,
	}

	out := SearchAddressesOut{
		Address: address,
		Err:     err,
	}

	return out
}

func (g *GeoGRPCClient) GeoCode(in GeoCodeIn) GeoCodeOut {
	res, err := g.client.GeoCode(context.Background(), &pb.GeoCodeRequest{Lat: in.Lat, Lng: in.Lng})

	out := GeoCodeOut{
		Lat: res.Lat,
		Lng: res.Lng,
		Err: err,
	}

	return out
}

func newClient(conf config.GeoRPC, protocol string) (*rpc.Client, error) {
	var (
		client *rpc.Client
		err    error
		host   = conf.Host
		port   = conf.Port
	)

	switch protocol {
	case rpcProtocol:
		client, err = rpc.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
		if err != nil {
			return nil, err
		}
		log.Println("rpc client connected")
		return client, nil

	case jsonrpcProtocol:
		// без этого костыля сервер редко успевает запуститься и коннект проваливается
		time.Sleep(1 * time.Second)

		client, err = jsonrpc.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
		if err != nil {
			return nil, err
		}
		log.Println("jsonrpc client connected")
		return client, nil

	default:
		return nil, fmt.Errorf("invalid protocol %s", protocol)
	}

}
