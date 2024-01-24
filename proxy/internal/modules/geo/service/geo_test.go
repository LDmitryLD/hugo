package service

import (
	"fmt"
	"projects/LDmitryLD/hugoproxy/proxy/internal/models"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/storage/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	query   = "москва"
	lat     = "55.7540471"
	lon     = "37.620405"
	address = models.Address{
		Lat: "10",
		Lon: "15",
	}
)

func TestGeo_SearchAddresses(t *testing.T) {
	adapterMock := mocks.NewGeoStorager(t)
	adapterMock.On("Select", query).Return(address, nil)

	geo := NewGeo(adapterMock)

	in := SearchAddressesIn{
		Query: query,
	}

	out := geo.SearchAddresses(in)

	assert.NotEmpty(t, out.Address)
	assert.Nil(t, out.Err)
}

func TestGeo_SearchAddresses2(t *testing.T) {
	adapterMock := mocks.NewGeoStorager(t)
	adapterMock.On("Select", query).Return(models.Address{}, fmt.Errorf("error"))
	adapterMock.On("Insert", query, lat, lon).Return(nil)

	geo := NewGeo(adapterMock)

	in := SearchAddressesIn{
		Query: query,
	}

	out := geo.SearchAddresses(in)

	assert.NotEmpty(t, out.Address)
	assert.Nil(t, out.Err)
}
