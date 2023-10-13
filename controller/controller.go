package controller

import (
	"aalsaedi265/slack_bot/database"
	"aalsaedi265/slack_bot/entity"
	
	"encoding/json"
	"io/ioutil"
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

	var student entity.Student
	database.Connector.First(&student, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}
func createStudent( w http.ResponseWriter, r *http.Request){
	
	requestBody, _ := ioutil.ReadAll(r.Body)
	var student entity.Student
	json.Unmarshal(requestBody, &student)

	

}