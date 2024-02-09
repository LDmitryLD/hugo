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

		r.Route("/debug/pprof", func(r chi.Router) {
			r.HandleFunc("/", pprof.Index)
			r.HandleFunc("/cmdline", pprof.Cmdline)
			r.HandleFunc("/profile", pprof.Profile)
			r.HandleFunc("/symbol", pprof.Symbol)
			r.HandleFunc("/trace", pprof.Trace)
			r.Handle("/allocs", pprof.Handler("allocs"))
			r.Handle("/block", pprof.Handler("block"))
			r.Handle("/goroutine", pprof.Handler("goroutine"))
			r.Handle("/mutex", pprof.Handler("mutex"))
			r.Handle("/heap", pprof.Handler("heap"))
			r.Handle("/threadcreate", pprof.Handler("threadcreate"))
		})
	})

	return r
}
