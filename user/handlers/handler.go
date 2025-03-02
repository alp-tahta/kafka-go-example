package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"user/models"
	"user/service"
)

type Handler struct {
	s service.ServiceInterface
}

func NewHandler(s service.ServiceInterface) *Handler {
	return &Handler{s: s}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) HealthChecker(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.s.CreateUser(user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
