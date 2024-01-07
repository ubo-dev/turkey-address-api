package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	log.Println("JSON API server running on port: ", s.listenAddr)
	router.HandleFunc("/cities", makeHttpHandleFunc(s.handleGetAllCities))

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleGetAllCities(w http.ResponseWriter, r *http.Request) error {
	cities, err := s.store.GetAllCities()
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
