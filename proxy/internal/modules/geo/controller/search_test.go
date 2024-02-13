package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"projects/LDmitryLD/hugoproxy/proxy/internal/infrastructure/responder"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/service"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules/geo/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeo_Search(t *testing.T) {
	geo := &GeoController{
		geo:       &service.Geo{},
		Responder: &responder.Respond{},
	}
	server := httptest.NewServer(http.HandlerFunc(geo.Search))
	r := SearchRequest{
		Query: "Москва",
	}

	body, _ := json.Marshal(r)

	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGeo_Search_Error(t *testing.T) {

	geoMock := mocks.NewGeorer(t)

	geoMock.On("SearchAddresses", service.SearchAddressesIn{Query: "BadQuery"}).Return(service.SearchAddressesOut{Err: errors.New("error")})

	geo := NewGeoController(geoMock)

	searchReq := SearchRequest{
		Query: "BadQuery",
	}

	reqBody, _ := json.Marshal(searchReq)

	s := httptest.NewServer(http.HandlerFunc(geo.Search))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
