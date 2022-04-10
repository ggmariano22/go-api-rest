package handlers

import (
	"encoding/json"
	"net/http"
	"go/go-lang-api/models"
	"go/go-lang-api/database"
	"github.com/gorilla/mux"
)

func Personalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func Personalidade(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Personalidade
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CreatePersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	database.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)
}

func DeletePersonalidade(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Personalidade
	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}

func EditPersonalidade(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Personalidade
	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}