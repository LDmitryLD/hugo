package modules

import (
	"projects/LDmitryLD/hugoproxy/proxy/internal/infrastructure/responder"
	authcontroller "projects/LDmitryLD/hugoproxy/proxy/internal/modules/auth/controller"
	geocontroller "projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/controller"
)

type Controllers struct {
	Auth authcontroller.Auther
	Geo  geocontroller.Georer
}

func NewControllers() *Controllers {
	responder := &responder.Respond{}
	authcontroller := authcontroller.NewAuth()
	geoController := geocontroller.NewGeoController(responder)

	return &Controllers{
		Auth: authcontroller,
		Geo:  geoController,
	}
}
