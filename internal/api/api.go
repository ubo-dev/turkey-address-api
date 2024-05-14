package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ubo-dev/turkey-address-api/internal/repository"
)

type APIServer struct {
	listenAddr string
	repository *repository.MysqlRepository
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func NewAPIServer(listenAddr string, repository *repository.MysqlRepository) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		repository: repository,
	}
}

func (s *APIServer) Run() {
	handler := http.NewServeMux()

	log.Println("JSON API server running on port: ", s.listenAddr)

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, World!"))
		if err != nil {
			panic(err)
		}
	})

	// cities : sehirler
	handler.HandleFunc("GET /city", makeHttpHandleFunc(s.handleGetAllCities))
	handler.HandleFunc("GET /city/{id}", makeHttpHandleFunc(s.handleGetCityById))

	// districts : ilceler
	handler.HandleFunc("GET /districts", makeHttpHandleFunc(s.handleGetAllDistricts))
	handler.HandleFunc("GET /district/{id}", makeHttpHandleFunc(s.handleGetDistrictById))
	handler.HandleFunc(
		"GET /district/getByCityId/{cityId}",
		makeHttpHandleFunc(s.handleGetDistrictByCityId),
	)

	// neighbourhoods : mahalleler
	handler.HandleFunc("GET /neighbourhood", makeHttpHandleFunc(s.handleGetAllNeighbourhoods))
	handler.HandleFunc(
		"GET /neighbourhood/getByZipCode/{zipCode}",
		makeHttpHandleFunc(s.handleGetGetNeighbourhoodsByZipCode),
	)
	handler.HandleFunc(
		"GET /neighbourhood/getByDistrictName/{districtName}",
		makeHttpHandleFunc(s.handleGetNeighbourhoodsByDistrictName),
	)
	handler.HandleFunc(
		"GET /neighbourhood/getByDistrictId/{id}",
		makeHttpHandleFunc(s.handleGetNeighbourhoodsByDistrictId),
	)

	err := http.ListenAndServe(s.listenAddr, handler)
	if err != nil {
		panic(err)
	}
}

func (s *APIServer) handleGetAllCities(w http.ResponseWriter, r *http.Request) error {
	cities, err := s.repository.GetAllCities()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, cities)
}

func (s *APIServer) handleGetCityById(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return err
	}

	city, err := s.repository.GetCityById(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, city)
}

func (s *APIServer) handleGetAllDistricts(w http.ResponseWriter, r *http.Request) error {
	districts, err := s.repository.GetAllDistricts()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, districts)
}

func (s *APIServer) handleGetDistrictById(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return err
	}

	districts, err := s.repository.GetDistrictById(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, districts)
}

func (s *APIServer) handleGetDistrictByCityId(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(r.PathValue("cityId"))
	if err != nil {
		return err
	}

	districts, err := s.repository.GetDistrictByCityId(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, districts)
}

func (s *APIServer) handleGetAllNeighbourhoods(w http.ResponseWriter, r *http.Request) error {
	neighbourhoods, err := s.repository.GetAllNeighbourhoods()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, neighbourhoods)
}

func (s *APIServer) handleGetGetNeighbourhoodsByZipCode(
	w http.ResponseWriter,
	r *http.Request,
) error {
	zipCode := r.PathValue("zipCode")

	neighbourhoods, err := s.repository.GetNeighbourhoodsByZipCode(zipCode)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, neighbourhoods)
}

func (s *APIServer) handleGetNeighbourhoodsByZipCode(
	w http.ResponseWriter,
	r *http.Request,
) error {
	zipCode := r.PathValue("zipCode")

	neighbourhoods, err := s.repository.GetNeighbourhoodsByZipCode(zipCode)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, neighbourhoods)
}

func (s *APIServer) handleGetNeighbourhoodsByDistrictName(
	w http.ResponseWriter,
	r *http.Request,
) error {
	districtName := r.PathValue("districtName")

	neighbourhoods, err := s.repository.GetNeighbourhoodsByDistrictName(districtName)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, neighbourhoods)
}

func (s *APIServer) handleGetNeighbourhoodsByDistrictId(
	w http.ResponseWriter,
	r *http.Request,
) error {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return err
	}

	neighbourhoods, err := s.repository.GetNeighbourhoodsByDistrictId(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, neighbourhoods)
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
