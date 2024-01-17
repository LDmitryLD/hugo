package responder

import (
	"encoding/json"
	"log"
	"net/http"
)

type Responder interface {
	OutputJSON(w http.ResponseWriter, responseData interface{})

	ErrorUnauthorized(w http.ResponseWriter, err error)
	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)
}

type Respond struct {
}

func (r *Respond) OutputJSON(w http.ResponseWriter, responseData interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		log.Println("Ошибка при отправке ответа:", err.Error())
		r.ErrorInternal(w, err)
	}
}

func (r *Respond) ErrorUnauthorized(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusUnauthorized)
}

func (r *Respond) ErrorBadRequest(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func (r *Respond) ErrorInternal(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
