package run

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"projects/LDmitryLD/hugoproxy/proxy/config"
	"projects/LDmitryLD/hugoproxy/proxy/internal/db"
	"projects/LDmitryLD/hugoproxy/proxy/internal/infrastructure/cache"
	"projects/LDmitryLD/hugoproxy/proxy/internal/infrastructure/router"
	"projects/LDmitryLD/hugoproxy/proxy/internal/infrastructure/server"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/service"
	"projects/LDmitryLD/hugoproxy/proxy/internal/storages"
	"projects/LDmitryLD/hugoproxy/proxy/rpc/geo"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	rpcProtocol     = "rpc"
	jsonrpcProtocol = "json-rpc"
)

type Application interface {
	Runner
	Bootstraper
}

type Runner interface {
	Run() error
}

type Bootstraper interface {
	Bootstrap(options ...interface{}) Runner
}

type App struct {
	conf     config.AppConf
	srv      server.Server
	rpc      server.Server
	Sig      chan os.Signal
	Storages *storages.Storages
	Services *modules.Services
}

func NewApp(conf config.AppConf) *App {
	return &App{conf: conf, Sig: make(chan os.Signal, 1)}
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())

	errGroup, ctx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		sigInt := <-a.Sig
		log.Println("signal interrupt revieved:", sigInt)
		cancel()
		return nil
	})

	errGroup.Go(func() error {
		err := a.srv.Serve(ctx)
		if err != nil && err != http.ErrServerClosed {
			log.Println("app: server error:", err)
			return err
		}
		return nil
	})

	if err := errGroup.Wait(); err != nil {
		return err
	}

	return nil
}

func (a *App) Bootstrap(options ...interface{}) Runner {

	_, sqlAdapter, err := db.NewSqlDB(a.conf.DB)
	if err != nil {
		log.Fatal("error init db: ", err)
	}

	cacheClient, err := cache.NewRedisClient(a.conf.Cache.Host, a.conf.Cache.Port)
	if err != nil {
		log.Fatal("error init cache:", err)
	}

	newStorages := storages.NewStorages(sqlAdapter, cacheClient)
	a.Storages = newStorages

	services := modules.NewSrvices(newStorages)
	a.Services = services

	geoRPC := geo.NewGeoServiceRPC(a.Services.Geo)
	RPCServer := rpc.NewServer()
	err = RPCServer.Register(geoRPC)
	if err != nil {
		log.Fatal("error init geo RPC:", err)
	}

	a.rpc, err = server.NewServerRPC(a.conf.RPCServer, RPCServer)
	if err != nil {
		log.Fatal("error new rpc server:", err)
	}
	go func() {
		err := a.rpc.Serve(context.Background())
		if err != nil {
			log.Fatal("app: server error", err)
		}
	}()

	client, err := newClient(a.conf.GeoRPC, a.conf.RPCServer.Type)
	if err != nil {
		log.Fatal("error init client:", err)
	}

	geoClientRPC := service.NewGeoRPC(client)

	controllers := modules.NewControllers(services, geoClientRPC)

	r := router.NewRouter(controllers)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.conf.Server.Port),
		Handler: r,
	}

	a.srv = server.NewHTTPServer(a.conf.Server, srv)

	return a
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
