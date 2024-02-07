package storage

import (
	"encoding/json"
	"log"
	"projects/LDmitryLD/hugoproxy/proxy/internal/models"
	"time"

	"github.com/go-redis/redis"
)

type GeoStorageProxy struct {
	storage *GeoStorage
	cache   *redis.Client
}

func NewGeoStorageProxy(storage *GeoStorage, cache *redis.Client) GeoStorager {
	return &GeoStorageProxy{
		storage: storage,
		cache:   cache,
	}
}

func (g *GeoStorageProxy) Select(query string) (models.Address, error) {
	startTime := time.Now()

	data, err := g.cache.Get(query).Result()

	duration := time.Since(startTime).Seconds()
	GeoControllerSearchCacheDuration.Observe(duration)

	if err == nil {
		var address models.Address

		err = json.Unmarshal([]byte(data), &address)
		if err == nil {
			log.Println("данные получены из кэша")
			return address, nil
		}
		log.Println("ошибка при разборе данных из кэша: ", err)
	}

	if err == redis.Nil {
		log.Println("данных нет в кэшэ")
	} else if err != nil {
		log.Println("ошибка при полученни данных из кэша: ", err)
	}
	log.Println("данные получены из бд")

	address, err := g.storage.Select(query)
	if err != nil {
		return models.Address{}, err
	}

	err = g.cache.Set(query, address, 5*time.Minute).Err()
	if err != nil {
		log.Println("ошибка при сохранении данных в кэш", err)
	} else {
		log.Println("данные записаны в кэш")
	}

	return address, nil
}

func (g *GeoStorageProxy) Insert(query string, lat string, lon string) error {
	address := models.Address{
		Lat: lat,
		Lon: lon,
	}

	err := g.cache.Set(query, address, 5*time.Minute).Err()
	if err != nil {
		log.Println("ошибка при записи данных кэш")
	} else {
		log.Println("данные записаны в кэш")
	}

	return g.storage.Insert(query, lat, lon)
}
