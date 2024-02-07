package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projects/LDmitryLD/hugoproxy/proxy/internal/infrastructure/responder"
	"projects/LDmitryLD/hugoproxy/proxy/internal/modules/auth/service"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	authRegisterRequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "auth_register_requests_total",
		Help: "Total number of requests",
	})
	authLoginRequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "auth_login_request_total",
		Help: "Total number of requests",
	})
	authRegisterDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "auth_register_duration_seconds",
		Help: "Reqeust duration in seconds",
	})
	authLoginDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "auth_login_duration_seconds",
		Help: "Request duration in seconds",
	})
)

func init() {
	prometheus.MustRegister(authRegisterRequestsTotal)
	prometheus.MustRegister(authLoginRequestsTotal)
	prometheus.MustRegister(authRegisterDuration)
	prometheus.MustRegister(authLoginDuration)
}

type Auther interface {
	Register(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
}

type Auth struct {
	auth service.Auther
	responder.Responder
}

func NewAuth(service service.Auther) Auther {
	return &Auth{
		auth:      service,
		Responder: &responder.Respond{},
	}
}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	authLoginRequestsTotal.Inc()

	var logReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&logReq); err != nil {
		log.Println("ошибка при декодировании запроса на вход", err.Error())
		a.Responder.ErrorBadRequest(w, err)
		return
	}

	out := a.auth.Login(r.Context(), service.AuthorizeIn{Name: logReq.Username, Password: logReq.Password})

	logResp := LoginResponse{
		Success: out.Success,
		Message: out.Message,
	}

	w.WriteHeader(http.StatusOK)
	a.OutputJSON(w, logResp)

	duration := time.Since(startTime).Seconds()
	authLoginDuration.Observe(duration)
}

func (a *Auth) Register(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	authRegisterRequestsTotal.Inc()

	var regReq RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&regReq); err != nil {
		log.Println("ошибка при декодировании запроса на регистрацию: ", err.Error())
		a.ErrorBadRequest(w, err)
		return
	}

	out := a.auth.Register(service.RegisterIn{Name: regReq.Username, Password: regReq.Password})

	if out.Error != nil {
		http.Error(w, out.Error.Error(), out.Status)
		return
	}

	regResp := RegisterReponse{
		Success: true,
		Message: fmt.Sprintf("Пользователь с именем %s зарегестрирован", regReq.Username),
	}

	w.WriteHeader(http.StatusOK)
	a.OutputJSON(w, regResp)

	duration := time.Since(startTime).Seconds()
	authRegisterDuration.Observe(duration)
}
