package run

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
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

	"golang.org/x/sync/errgroup"
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
	protocol := a.conf.RPCServer.Type

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

	a.rpc, err = server.GetServerRPC(a.conf.RPCServer, RPCServer, geo.NewGeoServiceGRPC(a.Services.Geo))
	if err != nil {
		log.Fatal("error init rpc server:", err)
	}
	go func() {
		err := a.rpc.Serve(context.Background())
		if err != nil {
			log.Fatal("app: server error", err)
		}
	}()

	geoClientRPC, err := service.GetlientRPC(protocol, a.conf.GeoRPC)
	if err != nil {
		log.Println("error init client:", err)
	}

	controllers := modules.NewControllers(services, geoClientRPC)

	r := router.NewRouter(controllers)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.conf.Server.Port),
		Handler: r,
	}

	a.srv = server.NewHTTPServer(a.conf.Server, srv)

	return a
}
