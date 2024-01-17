package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"projects/LDmitryLD/hugoproxy/proxy/internal/infrastructure/responder"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func Test_Geocode(t *testing.T) {
	geo := NewGeoController(&responder.Respond{})
	server := httptest.NewServer(http.HandlerFunc(geo.Geocode))
	r := GeocodeRequest{
		Lat: "12.123",
		Lng: "45.540",
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

func Test_APIHandler(t *testing.T) {
	geo := NewGeoController(&responder.Respond{})
	req, _ := http.NewRequest("GET", "/api/", nil)

	recorder := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/api/*", geo.ApiHandler)

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "Hello from API", recorder.Body.String())

}
