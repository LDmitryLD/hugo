package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

const (
	serverPort         = "SERVER_PORT"
	envShutdownTimeout = "SHUTDOWN_TIMEOUT"

	parseShutdownTimeoutError = "config: parse server shutdown timeout error"
)

type AppConf struct {
	Server    Server
	DB        DB
	Cache     Cache
	RPCServer RPCServer
	GeoRPC    GeoRPC
}

type DB struct {
	Driver   string
	Name     string
	User     string
	Password string
	Host     string
	Port     string
}

type Server struct {
	Port             string
	ShutdoundTimeout time.Duration
}

type RPCServer struct {
	Port          string
	ShutdoundTime time.Duration
	Type          string
}

type GeoRPC struct {
	Host string
	Port string
}

type Cache struct {
	Host string
	Port string
}

func NewAppConf() AppConf {
	port := os.Getenv(serverPort)

	return AppConf{
		Server: Server{
			Port: port,
		},
		DB: DB{
			Driver:   os.Getenv("DB_DRIVER"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
		},
		Cache: Cache{
			Host: os.Getenv("REDIS_HOST"),
			Port: os.Getenv("REDIS_PORT"),
		},
	}

}

func (a *AppConf) Init() {
	shutDownTimeOut, err := strconv.Atoi(os.Getenv(envShutdownTimeout))
	if err != nil {
		log.Fatal(parseShutdownTimeoutError)
	}
	shutDownTimeout := time.Duration(shutDownTimeOut) * time.Second

	a.Server.ShutdoundTimeout = shutDownTimeout

	a.RPCServer.Port = os.Getenv("RPC_PORT")
	a.RPCServer.Type = os.Getenv("RPC_PROTOCOL")
	a.GeoRPC.Host = os.Getenv("GEO_RPC_HOST")
	a.GeoRPC.Port = os.Getenv("GEO_RPC_PORT")

}
