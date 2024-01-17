package docs

import (
	acontroller "projects/LDmitryLD/hugoproxy/proxy/internal/modules/auth/controller"
)

// swagger:route POST /api/login auth LoginRequest
// Авторизация пользователя.
// responses:
//	200: LoginResponse

// swagger:parameters LoginRequest
type LoginRequest struct {
	// in:body
	Body acontroller.LoginRequest
}

// swagger:response LoginResponse
type LoginResponse struct {
	// in:body
	Body acontroller.LoginResponse
}

// swagger:route POST /api/register auth RegisterRequest
// Регистрация пользователя.
// responses:
//	200: RegisterReponse

// swagger:parameters RegisterRequest
type RegisterRequest struct {
	// in:body
	Body acontroller.RegisterRequest
}

// swagger:response RegisterReponse
type RegisterReponse struct {
	// in:body
	Body acontroller.RegisterReponse
}
