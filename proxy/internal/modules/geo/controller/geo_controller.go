package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"projects/LDmitryLD/hugoproxy/proxy/internal/infrastructure/responder"
	"projects/LDmitryLD/hugoproxy/proxy/internal/models"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/service"
)

type Georer interface {
	Geocode(http.ResponseWriter, *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	ApiHandler(w http.ResponseWriter, r *http.Request)
}

type GeoController struct {
	geo service.Georer
	responder.Responder
}

func NewGeoController(service service.Georer) Georer {
	return &GeoController{
		geo:       service,
		Responder: &responder.Respond{},
	}
}

func (g *GeoController) Geocode(w http.ResponseWriter, r *http.Request) {
	var geocodeRequest GeocodeRequest
	if err := json.NewDecoder(r.Body).Decode(&geocodeRequest); err != nil {
		log.Println("Ошибка при декодировании запроса: ", err)
		g.ErrorBadRequest(w, err)
		return
	}

	// var geocodeResponse GeocodeResponse
	// geocodeResponse.Addresses = []*Address{{Lat: geocodeRequest.Lat, Lon: geocodeRequest.Lng}}

	geocodeResponse := GeocodeResponse{
		Addresses: []*models.Address{{Lat: geocodeRequest.Lat, Lon: geocodeRequest.Lng}},
	}

	g.OutputJSON(w, geocodeResponse)
}

func (g *GeoController) Search(w http.ResponseWriter, r *http.Request) {
	var searchRequest SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&searchRequest); err != nil {
		log.Println("Ошибка при декодировании запроса: ", err)
		g.ErrorBadRequest(w, err)
		return
	}

	out := g.geo.SearchAddresses(service.SearchAddressesIn{Query: searchRequest.Query})
	if out.Err != nil {
		log.Println("Ошибка при получении адресса: ", out.Err)
		g.ErrorInternal(w, out.Err)
		return
	}

	searchResponse := SearchResponse{
		Addresses: []*models.Address{&out.Address},
	}

	g.OutputJSON(w, searchResponse)
}

func (g GeoController) ApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from API"))
}
