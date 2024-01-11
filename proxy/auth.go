package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("mysecret"), nil)
}

var users = make(map[string]string)

// swagger:route POST /api/login auth LoginRequest
// Авторизация пользователя.
// responses:
//	200: LoginResponse

// swagger:parameters LoginRequest
type LoginRequest struct {
	// in:body
	Username string `json:"username"`
	Password string `json:"password"`
}

// swagger:response LoginResponse
type LoginResponse struct {
	// in:body
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var logReq LoginRequest
	err := json.NewDecoder(r.Body).Decode(&logReq)
	if err != nil {
		log.Println("ошибка при декодировании запроса на вход", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pass, ok := users[logReq.Username]
	if !ok {
		log.Printf("ошибка: пользователь с именем %s не найден", logReq.Username)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Пользователь не найден"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(logReq.Password))
	if err != nil {
		log.Println("ошибка при сравнени паролей:", err.Error())
		http.Error(w, err.Error(), http.StatusOK)
		return
	}
	_, claims, _ := jwtauth.FromContext(r.Context())

	_, tokenString, _ := tokenAuth.Encode(claims)
	logResp := LoginResponse{
		Success: true,
		Message: tokenString,
	}
	json.NewEncoder(w).Encode(logResp)
	log.Println("LOGING")
}

// swagger:route POST /api/register auth RegisterRequest
// Регистрация пользователя.
// responses:
//	200: RegisterReponse

// swagger:parameters RegisterRequest
type RegisterRequest struct {
	// in:body
	Username string `json:"username"`
	Password string `json:"password"`
}

// swagger:response RegisterReponse
type RegisterReponse struct {
	// in:body
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var regReq RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&regReq)
	if err != nil {
		log.Println("ошибка при декодировании запроса на регистрацию: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(regReq.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("ошибка при генерации пароля: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
