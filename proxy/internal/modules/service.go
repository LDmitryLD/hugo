package modules

import (
	aservice "projects/LDmitryLD/hugoproxy/proxy/internal/modules/auth/service"
	geoservice "projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/service"
	"projects/LDmitryLD/hugoproxy/proxy/internal/storages"
)

type Services struct {
	Geo  geoservice.Georer
	Auth aservice.Auther
}

func NewSrvices(storages *storages.Storages) *Services {
	aService := aservice.NewAuth()
	geoService := geoservice.NewGeo(storages.Geo)

	return &Services{
		Geo:  geoService,
		Auth: aService,
	}
}
