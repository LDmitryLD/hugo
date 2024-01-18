package modules

import (
	authcontroller "projects/LDmitryLD/hugoproxy/proxy/internal/modules/auth/controller"
	geocontroller "projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/controller"
)

type Controllers struct {
	Auth authcontroller.Auther
	Geo  geocontroller.Georer
}

func NewControllers(services *Services) *Controllers {
	authcontroller := authcontroller.NewAuth(services.Auth)
	geoController := geocontroller.NewGeoController(services.Geo)

	return &Controllers{
		Auth: authcontroller,
		Geo:  geoController,
	}
}
