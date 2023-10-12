
package database

import (
	"log"
	"aalsaedi265/slack_bot/entity"

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

func Migrate(table *entity.Student){
	Connector.AutoMigrate(&table)
	log.Println("Migration completed")
}