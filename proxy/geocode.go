package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type GeocodeRequest struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type GeocodeResponse struct {
	Addresses []*Address `json:"addresses"`
}

func geocode(w http.ResponseWriter, r *http.Request) {
	var geocodeRequest GeocodeRequest
	err := json.NewDecoder(r.Body).Decode(&geocodeRequest)
	if err != nil {
		log.Println("Ошибка при декодировании запроса: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	req.Header.Set("Authorization", "Token 17ceca4516a27a2bc82188bbd4d524f1cec137a4")

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
