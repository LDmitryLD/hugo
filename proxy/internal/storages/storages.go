package storages

import (
	"projects/LDmitryLD/hugoproxy/proxy/internal/db/adapter"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/storage"

	"github.com/redis/go-redis/v9"
)

type Storages struct {
	Geo storage.GeoStorager
}

func NewStorages(sqlAdapter *adapter.SQLAdapter, cache *redis.Client) *Storages {
	geoStorage := storage.NewGeoStorage(sqlAdapter)
	proxy := storage.NewGeoStorageProxy(geoStorage, cache)

	return &Storages{
		Geo: proxy,
	}
}
