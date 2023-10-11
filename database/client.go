
package database

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	
	var err error
	Connector, err = gorm.Open("mssql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection Works Eurrieka")
	return nil
}