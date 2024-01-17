package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"projects/LDmitryLD/hugoproxy/proxy/internal/infrastructure/responder"

	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/client"
)

type Georer interface {
	Geocode(http.ResponseWriter, *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	ApiHandler(w http.ResponseWriter, r *http.Request)
}

type GeoController struct {
	responder responder.Responder
}

func NewGeoController(responder responder.Responder) Georer {
	return &GeoController{
		responder: responder,
	}
}

func (g *GeoController) Geocode(w http.ResponseWriter, r *http.Request) {
	var geocodeRequest GeocodeRequest
	if err := json.NewDecoder(r.Body).Decode(&geocodeRequest); err != nil {
		log.Println("Ошибка при декодировании запроса: ", err)
		g.responder.ErrorBadRequest(w, err)
		return
	}

	var geocodeResponse GeocodeResponse
	geocodeResponse.Addresses = []*Address{{Lat: geocodeRequest.Lat, Lon: geocodeRequest.Lng}}

	g.responder.OutputJSON(w, geocodeResponse)
}

func (g *GeoController) Search(w http.ResponseWriter, r *http.Request) {
	var searchRequest SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&searchRequest); err != nil {
		log.Println("Ошибка при декодировании запроса: ", err)
		g.responder.ErrorBadRequest(w, err)
		return
	}

	api := dadata.NewCleanApi(client.WithCredentialProvider(&client.Credentials{
		ApiKeyValue:    "d538755936a28def6bca48517dd287303cb0dae7",
		SecretKeyValue: "81081aa1fa5ca90caa8a69b14947b5876f58b8db",
	}))

	addresses, err := api.Address(context.Background(), searchRequest.Query)
	if err != nil {
		log.Println("Ошибка при получении адресса: ", err)
		g.responder.ErrorInternal(w, err)
		return
	}

	var searchResponse SearchResponse

	searchResponse.Addresses = []*Address{{Lat: addresses[0].GeoLat, Lon: addresses[0].GeoLon}}

	// if err := json.NewEncoder(w).Encode(&searchResponse); err != nil {
	// 	log.Println("Ошибка при кодировании ответа: ", err)
	// 	g.responder.ErrorInternal(w, err)
	// 	return
	// }

	g.responder.OutputJSON(w, searchResponse)
}

func (g GeoController) ApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from API"))
}
