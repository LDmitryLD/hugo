package geo

import "projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/service"

type GeoServiceRPC struct {
	geoService service.Georer
}

func NewGeoServiceRPC(geoService service.Georer) *GeoServiceRPC {
	return &GeoServiceRPC{geoService: geoService}
}

func (g *GeoServiceRPC) SearchAddresses(in service.SearchAddressesIn, out *service.SearchAddressesOut) error {
	*out = g.geoService.SearchAddresses(in)
	return nil
}

func (g *GeoServiceRPC) GeoCode(in service.GeoCodeIn, out *service.GeoCodeOut) error {
	*out = g.geoService.GeoCode(in)

	return nil
}
