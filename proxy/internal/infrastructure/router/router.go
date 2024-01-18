package router

import (
	"projects/LDmitryLD/hugoproxy/proxy/internal/infrastructure/middleware"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules"
	aservice "projects/LDmitryLD/hugoproxy/proxy/internal/modules/auth/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func NewRouter(controllers *modules.Controllers) *chi.Mux {
	r := chi.NewRouter()

	revProxy := middleware.NewReverseProxy("http://hugo", "1313")

	r.Use(revProxy.ReverseProxy)

	r.HandleFunc("/api/*", controllers.Geo.ApiHandler)

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
