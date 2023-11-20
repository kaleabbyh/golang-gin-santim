package main

import (
	"log"

	"gorm.io/gorm"

	"github.com/kaleabbyh/golang-santim/config"
	"github.com/kaleabbyh/golang-santim/models"
)

func MigrateTables(db *gorm.DB) error {
    err := db.AutoMigrate(models.User{}, 
						  models.Payment{},
		   				  models.Transaction{},
		  				  models.Account{},
						  models.Demo{},
						)
    if err != nil {
		log.Fatal("Error migrating tables:", err)
        return err
    }

    return nil
}

func main() {
	db, err := config.ConnectDB()
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
   
	MigrateTables(db)

}

