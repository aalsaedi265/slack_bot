
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
	database.Config()
}