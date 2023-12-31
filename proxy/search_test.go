package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(search))
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
