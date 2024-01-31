package storage

import (
	"encoding/json"
	"log"
	"projects/LDmitryLD/hugoproxy/proxy/internal/db/adapter"
	"projects/LDmitryLD/hugoproxy/proxy/internal/models"
	"time"

	"github.com/go-redis/redis"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=GeoStorager
type GeoStorager interface {
	Select(query string) (models.Address, error)
	Insert(query, lat, lon string) error
}

type GeoStorage struct {
	adapter adapter.SQLAdapterer
}

func NewGeoStorage(sqlAdapter adapter.SQLAdapterer) *GeoStorage {
	return &GeoStorage{
		adapter: sqlAdapter,
	}
}

func (g *GeoStorage) Select(query string) (models.Address, error) {
	return g.adapter.Select(query)
}

func (g *GeoStorage) Insert(query, lat, lon string) error {
	return g.adapter.Insert(query, lat, lon)
}

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
	data, err := g.cache.Get(query).Result()
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
	// var (
	// 	address models.Address
	// 	err     error
	// )

	// result, err := g.cache.Get(query).Result()
	// switch {
	// case err == redis.Nil:
	// 	log.Println("Данных нет в кэшэ, обращаемся к бд")
	// 	address, err = g.storage.Select(query)
	// 	if err != nil {
	// 		return address, err
	// 	}

	// 	err = g.cache.Set(query, address, 5*time.Minute).Err()
	// 	if err != nil {
	// 		log.Println("ошибка сохранения данных в кэш: ", err)
	// 		return address, nil
	// 	}
	// 	log.Println("данные записаны в кэш")
	// 	return address, nil

	// case err != nil:
	// 	log.Println("ошибка при получении данных из кэша: ", err)
	// 	address, err = g.storage.Select(query)
	// 	if err != nil {
	// 		return address, err
	// 	}

	// 	err = g.cache.Set(query, address, 5*time.Minute).Err()
	// 	if err != nil {
	// 		log.Println("ошибка сохранения данных в кэш: ", err)
	// 		return address, nil
	// 	}
	// 	log.Println("данные записаны в кэш")
	// 	return address, nil
	// }
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
