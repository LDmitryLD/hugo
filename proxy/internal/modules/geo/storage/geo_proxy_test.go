package storage

import (
	"context"
	"projects/LDmitryLD/hugoproxy/proxy/internal/db/adapter/mocks"
	"projects/LDmitryLD/hugoproxy/proxy/internal/models"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var (
	testLat   = "1"
	testLon   = "2"
	testQuery = "query"
	testValue = models.Address{
		Lat: testLat,
		Lon: testLon,
	}
)

func TestGeoStorageProxy_Select_FromCache(t *testing.T) {
	adapterMock := mocks.NewSQLAdapterer(t)
	cache := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer cache.Close()
	geoStorage := NewGeoStorage(adapterMock)

	proxy := NewGeoStorageProxy(geoStorage, cache)

	cache.Set(context.Background(), testQuery, testValue, 1*time.Minute)
	defer cache.Del(context.Background(), testQuery)

	got, err := proxy.Select(testQuery)

	assert.Equal(t, testValue, got)
	assert.Nil(t, err)
}

func TestGeoStorageProxy_Select_FromDB(t *testing.T) {
	adapterMock := mocks.NewSQLAdapterer(t)
	adapterMock.On("Select", testQuery).Return(testValue, nil)
	cache := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer cache.Close()
	geoStorage := NewGeoStorage(adapterMock)

	proxy := NewGeoStorageProxy(geoStorage, cache)

	got, err := proxy.Select(testQuery)

	assert.Equal(t, testValue, got)
	assert.Nil(t, err)
}

func TestGeoStprageProxy_Insert(t *testing.T) {
	adapterMock := mocks.NewSQLAdapterer(t)
	adapterMock.On("Insert", testQuery, testLat, testLon).Return(nil)
	cache := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer cache.Close()

	geoStorage := NewGeoStorage(adapterMock)
	proxy := NewGeoStorageProxy(geoStorage, cache)

	err := proxy.Insert(testQuery, testLat, testLon)

	assert.Nil(t, err)
}
