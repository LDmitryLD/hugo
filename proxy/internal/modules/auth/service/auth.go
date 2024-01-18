package service

import (
	"context"
	"errors"
	"net/http"
	"projects/LDmitryLD/hugoproxy/proxy/internal/infrastructure/tools/cryptography"

	"github.com/go-chi/jwtauth/v5"
)

var users = make(map[string]string)

var TokenAuth *jwtauth.JWTAuth

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("mysecret"), nil)
}

type Auth struct {
}

func NewAuth() Auther {
	return &Auth{}
}

func (a *Auth) Register(in RegisterIn) RegisterOut {
	hashPassword, err := cryptography.HashPassword(in.Password)
	if err != nil {
		return RegisterOut{
			Status: http.StatusInternalServerError,
			Error:  err,
		}
	}

	if _, ok := users[in.Name]; ok {
		return RegisterOut{
			Status: http.StatusBadRequest,
			Error:  errors.New("пользователь с таким именем уже зарегестрирован"),
		}
	}

	users[in.Name] = string(hashPassword)

	return RegisterOut{
		Status: http.StatusOK,
		Error:  nil,
	}
}

func (a *Auth) Login(ctx context.Context, in AuthorizeIn) AuthorizeOut {
	pass, ok := users[in.Name]
	if !ok {
		return AuthorizeOut{
			Success: false,
			Message: "Пользователь не найден",
		}
	}

	if !cryptography.CheckPassword(pass, in.Password) {
		return AuthorizeOut{
			Success: false,
			Message: "Неверный пароль",
		}
	}

	_, claims, _ := jwtauth.FromContext(ctx)

	_, tokenString, _ := TokenAuth.Encode(claims)

	return AuthorizeOut{
		Success: true,
		Message: tokenString,
	}
}
