package storage

import (
	"projects/LDmitryLD/hugoproxy/proxy/internal/db/adapter"
	"projects/LDmitryLD/hugoproxy/proxy/internal/models"
)

type GeoStorage struct {
	adapter adapter.SQLAdapter
}

func (g *GeoStorage) Select(query string) (models.Address, error) {
	return g.adapter.Select(query)
}

func (g *GeoStorage) Insert(query, lat, lon string) error {
	return g.adapter.Insert(query, lat, lon)
}
