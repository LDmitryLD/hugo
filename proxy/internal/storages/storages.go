package storages

import (
	"projects/LDmitryLD/hugoproxy/proxy/internal/db/adapter"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/storage"
)

type Storages struct {
	Geo storage.GeoStorager
}

func NewStorages(sqlAdapter *adapter.SQLAdapter) *Storages {
	return &Storages{
		Geo: storage.NewGeoStorage(sqlAdapter),
	}
}
