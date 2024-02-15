package modules

import (
	authcontroller "projects/LDmitryLD/hugoproxy/proxy/internal/modules/auth/controller"
	geocontroller "projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/controller"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/service"
)

type Controllers struct {
	Auth authcontroller.Auther
	Geo  geocontroller.Georer
}

func NewControllers(services *Services, geoRPC service.Georer) *Controllers {
	authcontroller := authcontroller.NewAuth(services.Auth)
	geoController := geocontroller.NewGeoController(geoRPC)

	return &Controllers{
		Auth: authcontroller,
		Geo:  geoController,
	}
}
