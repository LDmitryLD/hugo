package router

import (
	"net/http/pprof"
	"projects/LDmitryLD/hugoproxy/proxy/internal/infrastructure/middleware"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules"
	aservice "projects/LDmitryLD/hugoproxy/proxy/internal/modules/auth/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(controllers *modules.Controllers) *chi.Mux {
	r := chi.NewRouter()

	revProxy := middleware.NewReverseProxy("http://hugo", "1313")

	r.Use(revProxy.ReverseProxy)

	r.HandleFunc("/api/*", controllers.Geo.ApiHandler)
	r.Handle("/metrics", promhttp.Handler())

	r.Post("/api/login", controllers.Auth.Login)
	r.Post("/api/register", controllers.Auth.Register)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(aservice.TokenAuth))
		r.Use(jwtauth.Authenticator(aservice.TokenAuth))

		r.Post("/api/address/search", controllers.Geo.Search)
		r.Post("/api/address/geocode", controllers.Geo.Geocode)
	})

	return r
}

func NewPprof(controllers *modules.Controllers) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/debug/login", controllers.Auth.Login)
	r.Post("/debug/register", controllers.Auth.Register)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(aservice.TokenAuth))
		r.Use(jwtauth.Authenticator(aservice.TokenAuth))

		r.HandleFunc("/debug/pprof/", pprof.Index)
		r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		r.HandleFunc("/debug/pprof/profile", pprof.Profile)
		r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		r.HandleFunc("/debug/pprof/trace", pprof.Trace)
		r.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
		r.Handle("/debug/pprof/block", pprof.Handler("block"))
		r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
		r.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))
		r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
		r.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	})

	return r
}

// func NewPrometheus() *chi.Mux {
// 	r := chi.NewRouter()

// 	r.Handle("/metrics", promhttp.Handler())

// 	return r
// }
