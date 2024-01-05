package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/client"
)

type SearchRequest struct {
	Query string `json:"query"`
}

type SearchResponse struct {
	Addresses []*Address `json:"addresses"`
}

type Address struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func search(w http.ResponseWriter, r *http.Request) {
	var searchRequest SearchRequest
	err := json.NewDecoder(r.Body).Decode(&searchRequest)
	if err != nil {
		log.Println("Ошибка при декодировании запроса: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	api := dadata.NewCleanApi(client.WithCredentialProvider(&client.Credentials{
		ApiKeyValue:    "17ceca4516a27a2bc82188bbd4d524f1cec137a4",
		SecretKeyValue: "45710543548f02358064b56928a32789d2c71e7b",
	}))

	//log.Println("[QUERY] :", searchRequest.Query)

	addresses, err := api.Address(context.Background(), searchRequest.Query)
	if err != nil {
		log.Println("Ошибка при получении адресса: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//log.Println("[ADDRESSES] :", addresses[0])

	var searchResponse SearchResponse

	searchResponse.Addresses = []*Address{{Lat: addresses[0].GeoLat, Lon: addresses[0].GeoLon}}

	//log.Println("[SEARCHRESPONSE] : ", searchResponse.Addresses[0].Lat, searchResponse.Addresses[0].Lon)

	err = json.NewEncoder(w).Encode(&searchResponse)
	if err != nil {
		log.Println("Ошибка при кодировании ответа: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//log.Println("КОНЕЦ")
}
