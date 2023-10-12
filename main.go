
package main

import(
	"fmt"
	"net/http"
	// "os"
	// "time"
	"log"
	"aalsaedi265/slack_bot/database"
	"aalsaedi265/slack_bot/entity"


)
func main(){
	fmt.Println("the world")
	initDB()
	log.Println("Starting the HTTP server on port 8999")
	



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