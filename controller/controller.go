package controller

import (
	"aalsaedi265/slack_bot/database"
	"aalsaedi265/slack_bot/entity"
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
)

func GetAllStudent(w http.ResponseWriter, r *http.Request){
	var students []entity.Student
	database.Connector.Find(&students)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)
}

func GetStudentsByID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["Id"]
}