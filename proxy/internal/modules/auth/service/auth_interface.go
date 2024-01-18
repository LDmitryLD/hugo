package service

import (
	"context"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=Auther
type Auther interface {
	Register(in RegisterIn) RegisterOut
	Login(ctx context.Context, in AuthorizeIn) AuthorizeOut
}

type AuthorizeIn struct {
	Name     string
	Password string
}

type AuthorizeOut struct {
	Success bool
	Message string
}

type RegisterIn struct {
	Name     string
	Password string
}

type RegisterOut struct {
	Status int
	Error  error
}
