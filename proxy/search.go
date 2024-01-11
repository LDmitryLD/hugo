package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/client"
)

//swagger:route Post /api/address/search search SearchRequest
// Вычисление местанахождения по адрессу.
// security:
//   - Bearer: []
// responses:
//	200: SearchResponse
//

//swagger:parameters SearchRequest
type SearchRequest struct {
	//Qury - запрос, представляющий собой адрес
	//in: body
	Query string `json:"query"`
}

//swagger:response SearchResponse
type SearchResponse struct {
	// Addresses содержит список адрессов
	// in: body
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
		ApiKeyValue:    "d538755936a28def6bca48517dd287303cb0dae7",
		SecretKeyValue: "81081aa1fa5ca90caa8a69b14947b5876f58b8db",
	}))

	addresses, err := api.Address(context.Background(), searchRequest.Query)
	if err != nil {
		log.Println("Ошибка при получении адресса: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var searchResponse SearchResponse

	searchResponse.Addresses = []*Address{{Lat: addresses[0].GeoLat, Lon: addresses[0].GeoLon}}

	err = json.NewEncoder(w).Encode(&searchResponse)
	if err != nil {
		log.Println("Ошибка при кодировании ответа: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
