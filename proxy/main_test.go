package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func Test_apiHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/", nil)

	recorder := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/api/*", apiHandler)

	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "Hello from API", recorder.Body.String())
}

func Test_Proxy(t *testing.T) {
	r := chi.NewRouter()
	r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test hugo"))
	})

	testServer := httptest.NewServer(r)
	url, _ := url.Parse(testServer.URL)

	proxy := NewReverseProxy("http://localhost", url.Port())

	router := chi.NewRouter()
	router.Use(proxy.ReverseProxy)
	router.HandleFunc("/api/*", apiHandler)
	server := httptest.NewServer(router)

	resp, err := http.Get(server.URL + "/api/test")
	if err != nil {
		t.Error("ошибка при запросе на /api/test:", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("ошибка при чтении тела ответа с /api/test:", err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Hello from API", string(body))

	resp, err = http.Get(server.URL + "/test")
	if err != nil {
		t.Error("ошибка при запросе на /test:", err)
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Error("ошибка при чтении тела ответа с /test:", err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Test hugo", string(body))
}
