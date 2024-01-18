package service

import (
	"github.com/ekomobile/dadata/v2/api/model"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=Georer
type Georer interface {
	SearchAddresses(in SearchAddressesIn) SearchAddressesOut
}

type SearchAddressesIn struct {
	Query string
}

type SearchAddressesOut struct {
	Addresses []*model.Address
	Err       error
}
