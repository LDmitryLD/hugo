package service

import (
	"context"
	"net/http"
	"projects/LDmitryLD/hugoproxy/proxy/internal/infrastructure/tools/cryptography"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAuth(t *testing.T) {
	expect := &Auth{}

	got := NewAuth()

	assert.Equal(t, expect, got)
}

func TestAuth_Register_Success(t *testing.T) {
	auth := NewAuth()

	in := RegisterIn{
		Name:     "testuser",
		Password: "123",
	}

	out := auth.Register(in)

	assert.Nil(t, out.Error)
	assert.Equal(t, http.StatusOK, out.Status)
}

func TestAuth_Register_AlreadyExists(t *testing.T) {
	auth := NewAuth()

	users["testuser"] = "123"

	in := RegisterIn{
		Name:     "testuser",
		Password: "123",
	}

	out := auth.Register(in)

	assert.NotNil(t, out.Error)
	assert.Equal(t, http.StatusBadRequest, out.Status)
}

func TestAuth_Login_Success(t *testing.T) {
	auth := NewAuth()

	hashPassword, _ := cryptography.HashPassword("123")
	users["testuser"] = string(hashPassword)

	ctx := context.Background()
	in := AuthorizeIn{
		Name:     "testuser",
		Password: "123",
	}

	out := auth.Login(ctx, in)

	assert.True(t, out.Success)
	assert.NotEmpty(t, out.Message)
}

func TestAuth_Login_NotFound(t *testing.T) {
	auth := NewAuth()

	ctx := context.Background()
	in := AuthorizeIn{
		Name:     "user",
		Password: "123",
	}

	out := auth.Login(ctx, in)

	assert.False(t, out.Success)
	assert.Equal(t, "Пользователь не найден", out.Message)
}

func TestAuth_Login_WrondPassword(t *testing.T) {
	auth := NewAuth()

	users["testuser"] = "password"

	ctx := context.Background()
	in := AuthorizeIn{
		Name:     "testuser",
		Password: "wrondpassword",
	}

	out := auth.Login(ctx, in)

	assert.False(t, out.Success)
	assert.Equal(t, "Неверный пароль", out.Message)
}
