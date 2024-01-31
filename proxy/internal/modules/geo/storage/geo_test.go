package storage

import (
	"projects/LDmitryLD/hugoproxy/proxy/internal/db/adapter/mocks"
	"projects/LDmitryLD/hugoproxy/proxy/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	query       = "test"
	lat         = "10"
	lon         = "15"
	testAddress = models.Address{
		Lat: "10",
		Lon: "15",
	}
)

func TestGeoStorage_Select(t *testing.T) {
	adapterMock := mocks.NewSQLAdapterer(t)
	adapterMock.On("Select", query).Return(testAddress, nil)

	geo := NewGeoStorage(adapterMock)

	res, err := geo.Select(query)

	assert.Nil(t, err)
	assert.Equal(t, testAddress, res)
}

func TestGeoStorage_Insert(t *testing.T) {
	adapterMock := mocks.NewSQLAdapterer(t)
	adapterMock.On("Insert", query, lat, lon).Return(nil)

	geo := NewGeoStorage(adapterMock)

	err := geo.Insert(query, lat, lon)

	assert.Nil(t, err)
}
