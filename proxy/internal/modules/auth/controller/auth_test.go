package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules/auth/service"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules/auth/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuth_Login_BadRequest(t *testing.T) {

	auth := NewAuth(service.NewAuth())

	req := map[string]interface{}{"username": 123, "password": "password"}
	reqJSON, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(auth.Login))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal("ошибка при выполнении тестового запроса:", err.Error())
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestAuth_Login(t *testing.T) {
	authMock := mocks.NewAuther(t)

	authMock.On("Login", mock.Anything, service.AuthorizeIn{Name: "test", Password: "123"}).Return(service.AuthorizeOut{Success: true, Message: "token"})

	auth := NewAuth(authMock)

	logindReq := LoginRequest{
		Username: "test",
		Password: "123",
	}

	reqBody, _ := json.Marshal(logindReq)

	s := httptest.NewServer(http.HandlerFunc(auth.Login))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var logResp LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&logResp)
	if err != nil {
		t.Fatal("ошибка при декодировании тестового ответа:", err.Error())
	}

	assert.True(t, logResp.Success)
	assert.NotEmpty(t, logResp.Message)
}

func TestAuth_Register_BadRequest(t *testing.T) {
	auth := NewAuth(service.NewAuth())

	req := map[string]interface{}{"username": 123, "password": "password"}
	reqJSON, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(auth.Register))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal("ошибка при выполнении тестового запроса:", err.Error())
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestAuth_Register(t *testing.T) {
	authMock := mocks.NewAuther(t)

	authMock.On("Register", service.RegisterIn{Name: "test", Password: "123"}).Return(service.RegisterOut{Status: http.StatusOK, Error: nil})

	auth := NewAuth(authMock)

	req := RegisterRequest{
		Username: "test",
		Password: "123",
	}

	reqBody, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(auth.Register))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var regResp RegisterReponse
	err = json.NewDecoder(resp.Body).Decode(&regResp)
	if err != nil {
		t.Fatal("ошибка при декодировании тестового ответа:", err.Error())
	}

	assert.True(t, regResp.Success)
	assert.Equal(t, "Пользователь с именем test зарегестрирован", regResp.Message)
}

func TestAtuh_Register_Error(t *testing.T) {
	authMock := mocks.NewAuther(t)

	authMock.On("Register", service.RegisterIn{Name: "test", Password: "123"}).Return(service.RegisterOut{Status: http.StatusBadRequest, Error: errors.New("пользователь с таким именем уже зарегестрирован")})

	auth := NewAuth(authMock)

	req := RegisterRequest{
		Username: "test",
		Password: "123",
	}

	reqBody, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(auth.Register))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
