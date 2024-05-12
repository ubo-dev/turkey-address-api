package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ubo-dev/turkey-address-api/internal/repository"
)

type APIServer struct {
	listenAddr string
	repository repository.Repository
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func NewAPIServer(listenAddr string, repository repository.Repository) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		repository: repository,
	}
}

func (s *APIServer) Run() {

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.HandleFunc("GET /cities", makeHttpHandleFunc(s.handleGetAllCities))
}

func (s *APIServer) handleGetAllCities(w http.ResponseWriter, r *http.Request) error {
	cities, err := s.repository.GetAllCities()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, cities)
}

func makeHttpHandleFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
