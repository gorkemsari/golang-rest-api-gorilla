package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorkemsari/golang-rest-api/helper"
	repo "github.com/gorkemsari/golang-rest-api/repository"
)

func CityAll(w http.ResponseWriter, r *http.Request) {

	cities, err := repo.GetAllCities()
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, cities)
}

func City(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	cities, err := repo.GetCityById(id)
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, cities)
}
