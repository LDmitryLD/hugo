package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"projects/LDmitryLD/hugoproxy/proxy/mermaid"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	// go WorkerTest()

	r := chi.NewRouter()

	revProxy := NewReverseProxy("http://hugo", "1313")

	r.Use(revProxy.ReverseProxy)

	r.HandleFunc("/api/*", apiHandler)

	r.Post("/api/address/search", search)
	r.Post("/api/address/geocode", geocode)

	http.ListenAndServe(":8080", r)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from API"))
}

type ReverseProxy struct {
	host string
	port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/api/") {

			url, err := url.Parse(rp.host + ":" + rp.port)
			if err != nil {
				log.Fatal(err)
			}
			proxy := httputil.NewSingleHostReverseProxy(url)
			// r.URL.Scheme = "http"
			// r.URL.Host = "hugo:1313"
			r.Host = "hugo:1313"

			proxy.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}

	})

}

const content = `---
menu:
    before:
        name: tasks
        weight: 5
title: Обновление данных в реальном времени
---

# Задача: Обновление данных в реальном времени

Напишите воркер, который будет обновлять данные в реальном времени, на текущей странице.
Текст данной задачи менять нельзя, только время и счетчик.

Файл данной страницы: /app/static/tasks/_index.md 

Должен меняться счетчик и время:

Текущее время: %s

Счетчик: %d



## Критерии приемки:
- [ ] Воркер должен обновлять данные каждые 5 секунд
- [ ] Счетчик должен увеличиваться на 1 каждые 5 секунд
- [ ] Время должно обновляться каждые 5 секунд
`

func WorkerTest() {
	t := time.NewTicker(5 * time.Second)
	rand.Seed(time.Now().UnixNano())
	var tree *mermaid.AVLTree
	var treeCountre int
	var b byte = 0
	for {
		select {
		case <-t.C:
			graph := mermaid.MakeGraph()

			if treeCountre == 0 || treeCountre == 100 {
				tree = mermaid.GenerateTree(5)
				treeCountre = 5
			} else {
				tree.Insert(rand.Intn(150))
				treeCountre++
			}

			mer := tree.ToMermaid()
			err := os.WriteFile("/app/static/_index.md", []byte(fmt.Sprintf(content, time.Now().Format("2006-01-02 15:04:05"), b)), 0644)
			if err != nil {
				log.Println("ошибка при записи файла:", err)
			}
			b++

			err = os.WriteFile("/app/static/tasks/graph.md", []byte(fmt.Sprintf(mermaid.GraphConten, graph, graph)), 0644)
			if err != nil {
				log.Println("ошибка при записи файла:", err)
			}

			err = os.WriteFile("/app/static/tasks/binary.md", []byte(fmt.Sprintf(mermaid.BinaryContent, mer, mer)), 0644)
			if err != nil {
				log.Println("ошибка при записи файла:", err)
			}
		default:
			_ = 4
		}
	}
}
