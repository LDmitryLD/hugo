package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

//swagger:route Post /api/address/geocode geocode GeocodeRequest
// Вычисление адресса по широте и долготе.
// security:
//   - Bearer: []
//
// Response:
//	200: GeocodeResponse

//swagger:parameters GeocodeRequest
type GeocodeRequest struct {
	// Lat - широта
	// in: body
	// required: true
	Lat string `json:"lat"`
	// Lng - долгота
	// in: body
	// required: true
	Lng string `json:"lng"`
}

//swagger:response GeocodeResponse
type GeocodeResponse struct {
	// in: body
	// Addresses содержит список адрессов
	Addresses []*Address `json:"addresses"`
}

func geocode(w http.ResponseWriter, r *http.Request) {
	var geocodeRequest GeocodeRequest
	err := json.NewDecoder(r.Body).Decode(&geocodeRequest)
	if err != nil {
		log.Println("Ошибка при декодировании запроса: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, _ := json.Marshal(geocodeRequest)
	client := http.DefaultClient
	req, err := http.NewRequest("POST", "http://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", bytes.NewBuffer(data))
	if err != nil {
		log.Println("Ошибка при создании запроса: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token d538755936a28def6bca48517dd287303cb0dae7")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Ошибка при отправке запроса: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var suggestions Suggestions
	err = json.NewDecoder(resp.Body).Decode(&suggestions)
	if err != nil {
		log.Println("Ошибка декодировании ответа", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var geocodeResponse GeocodeResponse
	geocodeResponse.Addresses = []*Address{{Lat: geocodeRequest.Lat, Lon: geocodeRequest.Lng}}

	err = json.NewEncoder(w).Encode(&geocodeResponse)
	if err != nil {
		log.Println("Ошибка при отправке ответа: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type Suggestions struct {
	Suggestions []Suggestion `json:"suggestions"`
}

type Suggestion struct {
	Value             string `json:"value"`
	UnrestrictedValue string `json:"unrestricted_value"`
	Data              Data   `json:"data"`
}

type Data struct {
}
