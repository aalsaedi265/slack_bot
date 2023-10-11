
package main

import(
	"fmt"
	// "os"
	// "time"
	"aalsaedi265/slack_bot/database"


)
func main(){
	fmt.Println("the world")
	initDB()

}

func initDB(){
	config := 
				database.Config{
					ServerName: "localhost:3306",
					User:       "root",
					Password:   "1234",
					DB:         "Student_Details",
				}
	connectingString = database.GetConnectionString(config)
	
}