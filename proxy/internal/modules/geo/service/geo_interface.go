package service

import "projects/LDmitryLD/hugoproxy/proxy/internal/models"

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=Georer
type Georer interface {
	SearchAddresses(in SearchAddressesIn) SearchAddressesOut
}

type SearchAddressesIn struct {
	Query string
}

type SearchAddressesOut struct {
	Address models.Address
	Err     error
}
