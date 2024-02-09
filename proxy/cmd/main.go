package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"projects/LDmitryLD/hugoproxy/proxy/config"
	"projects/LDmitryLD/hugoproxy/proxy/internal/db"
	"projects/LDmitryLD/hugoproxy/proxy/internal/infrastructure/router"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules"
	"projects/LDmitryLD/hugoproxy/proxy/internal/storages"
	"syscall"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

func main() {

	confDB := config.NewAppConf().DB
	_, sqlAdapter, err := db.NewSqlDB(confDB)
	if err != nil {
		log.Fatal("ошибка при инициализации БД:", err)
	}

	cach := redis.NewClient(&redis.Options{
		Addr: "localhost:6378",
	})

	pong, err := cach.Ping().Result()
	if err != nil {
		fmt.Println("ошибка соединения с Redis:", err)
	}
	fmt.Println("соединение с Redis успешно:", pong)

	storages := storages.NewStorages(sqlAdapter, cach)

	services := modules.NewSrvices(storages)

	controllers := modules.NewControllers(services)

	r := router.NewRouter(controllers)

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Starting server")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error: ", err.Error())
		}
	}()

	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server stopped")
}
