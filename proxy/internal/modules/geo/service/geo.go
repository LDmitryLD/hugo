package service

import (
	"context"
	"log"
	"projects/LDmitryLD/hugoproxy/proxy/internal/models"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/storage"

	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/client"
)

type Geo struct {
	storage storage.GeoStorage
}

func NewGeo() Georer {
	return &Geo{}
}

func (g *Geo) SearchAddresses(in SearchAddressesIn) SearchAddressesOut {

	address, err := g.storage.Select(in.Query)
	if err != nil {
		res, err := searchFromAPI(in.Query)
		if err != nil {
			return SearchAddressesOut{
				Err: err,
			}
		}

		if err = g.storage.Insert(in.Query, res.Lat, res.Lon); err != nil {
			log.Println("ошибка при кэшировании данных:", err.Error())
		}

		log.Println("Данные закэшированны")

		return SearchAddressesOut{
			Address: res,
		}
	}

	log.Println("исползуем данные из кэша")
	return SearchAddressesOut{
		Address: address,
		Err:     nil,
	}
}

func searchFromAPI(query string) (models.Address, error) {
	api := dadata.NewCleanApi(client.WithCredentialProvider(&client.Credentials{
		ApiKeyValue:    "d538755936a28def6bca48517dd287303cb0dae7",
		SecretKeyValue: "81081aa1fa5ca90caa8a69b14947b5876f58b8db",
	}))

	addresses, err := api.Address(context.Background(), query)
	if err != nil {
		return models.Address{}, err
	}

	res := models.Address{
		Lat: addresses[0].GeoLat,
		Lon: addresses[0].GeoLon,
	}

	return res, nil
}
