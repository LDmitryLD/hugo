package main

import (
	"log"
	_ "net/http/pprof"
	"os"
	"projects/LDmitryLD/hugoproxy/proxy/config"
	"projects/LDmitryLD/hugoproxy/proxy/run"

	_ "github.com/lib/pq"
)

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println("ошибка при чтени .env файла:", err)
	// }

	conf := config.NewAppConf()

	conf.Init()

	app := run.NewApp(conf)

	if err := app.Bootstrap().Run(); err != nil {
		log.Printf("error: %s", err.Error())
		os.Exit(2)
	}

	// confDB := config.NewAppConf().DB
	// _, sqlAdapter, err := db.NewSqlDB(confDB)
	// if err != nil {
	// 	log.Fatal("ошибка при инициализации БД:", err)
	// }

	// cach := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6378",
	// })

	// pong, err := cach.Ping().Result()
	// if err != nil {
	// 	fmt.Println("ошибка соединения с Redis:", err)
	// }
	// fmt.Println("соединение с Redis успешно:", pong)

	// storages := storages.NewStorages(sqlAdapter, cach)

	// services := modules.NewSrvices(storages)

	// controllers := modules.NewControllers(services)

	// r := router.NewRouter(controllers)

	// s := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: r,
	// }

	// sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// go func() {
	// 	log.Println("Starting server")
	// 	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatal("Server error: ", err.Error())
	// 	}
	// }()

	// <-sigChan

	// ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer cancel()

	// if err := s.Shutdown(ctx); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Server stopped")
}
