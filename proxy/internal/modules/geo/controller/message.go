package controller

type GeocodeRequest struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type GeocodeResponse struct {
	Addresses []*Address `json:"addresses"`
}

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
