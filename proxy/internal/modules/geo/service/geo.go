package service

import (
	"context"

	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/client"
)

type Geo struct {
}

func NewGeo() Georer {
	return &Geo{}
}

func (g *Geo) SearchAddresses(in SearchAddressesIn) SearchAddressesOut {

	api := dadata.NewCleanApi(client.WithCredentialProvider(&client.Credentials{
		ApiKeyValue:    "d538755936a28def6bca48517dd287303cb0dae7",
		SecretKeyValue: "81081aa1fa5ca90caa8a69b14947b5876f58b8db",
	}))

	addresses, err := api.Address(context.Background(), in.Query)
	if err != nil {
		return SearchAddressesOut{
			Err: err,
		}
	}

	return SearchAddressesOut{
		Addresses: addresses,
		Err:       nil,
	}

}
