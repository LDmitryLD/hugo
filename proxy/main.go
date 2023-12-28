package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	revProxy := NewReverseProxy("http://hugo", "1313")
	r.Use(revProxy.ReverseProxy)
	r.HandleFunc("/api/*", apiHandler)

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

// const content = ``

// func WorkerTest() {
// 	t := time.NewTicker(1 * time.Second)
// 	var b byte = 0
// 	for {
// 		select {
// 		case <-t.C:
// 			err := os.WriteFile("/app/static/_index.md", []byte(fmt.Sprintf(content, b)), 0644)
// 			if err != nil {
// 				log.Println(err)
// 			}
// 			b++
// 		}
// 	}
// }
