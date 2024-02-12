package service

import (
	"net/rpc"
)

// type GeoProvider interface {
// 	AddressSearch(input string) ([]*models.Address, error)
// 	GeoCode(lat, lng string) ([]*models.Address, error)
// }

type GeoRPC struct {
	client *rpc.Client
}

func NewGeoRPC(client *rpc.Client) *GeoRPC {
	return &GeoRPC{
		client: client,
	}
}

func (g *GeoRPC) SearchAddresses(in SearchAddressesIn) SearchAddressesOut {
	var out SearchAddressesOut
	err := g.client.Call("GeoServiceRPC.SearchAddresses", in, &out)
	if err != nil {
		out.Err = err
	}

	return out
}

func (g *GeoRPC) GeoCode(in GeoCodeIn) GeoCodeOut {
	var out GeoCodeOut
	err := g.client.Call("GeoServiceRPC.GeoCode", in, &out)
	if err != nil {
		out.Err = err
	}

	return out
}
