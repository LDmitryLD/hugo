package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginHandler(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(loginHandler))
	defer s.Close()

	pass, _ := bcrypt.GenerateFromPassword([]byte("123test"), bcrypt.DefaultCost)
	users["test"] = string(pass)

	logReq := LoginRequest{
		Username: "test",
		Password: "123test",
	}

	body, err := json.Marshal(logReq)
	if err != nil {
		t.Fatal("ошибка при кодировании тестового запроса:", err.Error())
	}

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal("ошибка при выполнении тестового запроса:", err.Error())
	}
	defer resp.Body.Close()

	var logResp LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&logResp)
	if err != nil {
		t.Fatal("ошибка при декодировании тестового ответа:", err.Error())
	}

	assert.Equal(t, true, logResp.Success)
	assert.NotEqual(t, "", logResp.Message)
}

func TestLoginHandler_NotFound(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(loginHandler))
	defer s.Close()

	logReq := LoginRequest{
		Username: "test2",
		Password: "321test",
	}

	body, err := json.Marshal(logReq)
	if err != nil {
		t.Fatal("ошибка при кодировании тестового запроса:", err.Error())
	}

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal("ошибка при выполнении тестового запроса:", err.Error())
	}
	defer resp.Body.Close()
	bodyS, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Пользователь не найден", string(bodyS))
}

func TestLoginHandler_BadRequest(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(loginHandler))
	defer s.Close()

	req := "{ ID: 123}"

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		t.Fatal("ошибка при выполнении тестового запроса:", err.Error())
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestRegisterHandler(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(registerHandler))
	defer s.Close()

	regReq := RegisterRequest{
		Username: "test",
		Password: "123test",
	}

	reqBody, err := json.Marshal(regReq)
	if err != nil {
		t.Fatal("ошибка при кодировании тестового запроса:", err.Error())
	}

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal("ошибка при выполнении тестового запроса:", err.Error())
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var regResp RegisterReponse
	err = json.NewDecoder(resp.Body).Decode(&regResp)
	if err != nil {
		t.Fatal("ошибка при кодировании тестового ответа:", err.Error())
	}

	assert.Equal(t, true, regResp.Success)
	assert.Equal(t, "Пользователь с именем test зарегестрирован", regResp.Message)
}

func TestRegisterHandler_BadRequest(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(registerHandler))
	defer s.Close()

	req := "{ ID: 123}"

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer([]byte(req)))
	if err != nil {
		t.Fatal("ошибка при выполнении тестового запроса:", err.Error())
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

}
