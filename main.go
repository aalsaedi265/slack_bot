package main

import (
	"fmt"
	"net/http"

	// "os"
	// "time"
	"aalsaedi265/slack_bot/database"
	"aalsaedi265/slack_bot/entity"
	"aalsaedi265/slack_bot/controller"
	"log"

	"github.com/gorilla/mux"
)
func main(){

	fmt.Println("the world")
	initDB()
	log.Println("Starting the HTTP server on port 8999")
	router := mux.NewRouter().StrictSlash(true)
	initaliseHanlders(router)
	log.Fatal(http.ListenAndServe(":8999", router))
}

func initaliseHanlders(router *mux.Router){
	router.HandleFunc("/get", controller.GetAllStudent).Methods("GET")
}

func initDB(){
	config := 
				database.Config{
					ServerName: "localhost:3306",
					User:       "root",
					Password:   "1234",
					DB:         "Student_Details",
				}
	connectingString := database.GetConnectionString(config)
	err := database.Connect(connectingString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Student{})	
}