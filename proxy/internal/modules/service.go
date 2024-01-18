package modules

import (
	aservice "projects/LDmitryLD/hugoproxy/proxy/internal/modules/auth/service"
	geoservice "projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/service"
)

type Services struct {
	Geo  geoservice.Georer
	Auth aservice.Auther
}

func NewSrvices() *Services {
	return &Services{
		Geo:  geoservice.NewGeo(),
		Auth: aservice.NewAuth(),
	}
}
