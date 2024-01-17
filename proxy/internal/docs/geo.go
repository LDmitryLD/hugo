package docs

import (
	geocontroller "projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/controller"
)

//swagger:route Post /api/address/geocode geocode GeocodeRequest
// Вычисление адресса по широте и долготе.
// security:
//   - Bearer: []
// responses:
//	200: GeocodeResponse

//swagger:parameters GeocodeRequest
type GeocodeRequest struct {
	// Lat - широта
	// Lng - долгота
	// in: body
	// required: true
	Body geocontroller.GeocodeRequest
}

//swagger:response GeocodeResponse
type GeocodeResponse struct {
	// in: body
	Body geocontroller.GeocodeResponse
}

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
	Body geocontroller.SearchRequest
}

//swagger:response SearchResponse
type SearchResponse struct {
	// Addresses содержит список адрессов
	// in: body
	Body geocontroller.SearchResponse
}
