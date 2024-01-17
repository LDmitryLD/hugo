package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projects/LDmitryLD/hugoproxy/proxy/internal/infrastructure/responder"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

type Auther interface {
	Register(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
}

type Auth struct {
	responder.Responder
}

func NewAuth() Auther {
	return &Auth{
		Responder: &responder.Respond{},
	}
}

var TokenAuth *jwtauth.JWTAuth

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("mysecret"), nil)
}

var users = make(map[string]string)

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	var logReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&logReq); err != nil {
		log.Println("ошибка при декодировании запроса на вход", err.Error())
		a.Responder.ErrorBadRequest(w, err)
		return
	}

	pass, ok := users[logReq.Username]
	if !ok {
		log.Printf("ошибка: пользователь с именем %s не найден", logReq.Username)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Пользователь не найден"))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(logReq.Password)); err != nil {
		log.Println("ошибка при сравнени паролей:", err.Error())
		http.Error(w, err.Error(), http.StatusOK)
		return
	}

	_, claims, _ := jwtauth.FromContext(r.Context())

	_, tokenString, _ := TokenAuth.Encode(claims)
	logResp := LoginResponse{
		Success: true,
		Message: tokenString,
	}
	json.NewEncoder(w).Encode(logResp)
	log.Println("LOGING")
}

func (a *Auth) Register(w http.ResponseWriter, r *http.Request) {
	var regReq RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&regReq); err != nil {
		log.Println("ошибка при декодировании запроса на регистрацию: ", err.Error())
		a.Responder.ErrorBadRequest(w, err)
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(regReq.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("ошибка при генерации пароля: ", err.Error())
		a.Responder.ErrorInternal(w, err)
		return
	}

	users[regReq.Username] = string(pass)
	regResp := RegisterReponse{
		Success: true,
		Message: fmt.Sprintf("Пользователь с именем %s зарегестрирован", regReq.Username),
	}

	json.NewEncoder(w).Encode(regResp)
	log.Println("REGISTERED")
}
