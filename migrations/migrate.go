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
						)
    if err != nil {
        return err
    }

    return nil
}

func main() {
	db, err := config.ConnectDB()
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
   
	err = MigrateTables(db)
    if err != nil {
        log.Fatal("Error migrating tables:", err)
    }


}

	// err = CreateUserTable(db)
	// if err != nil {
	// 	log.Fatal("Error  creating user table:", err)
		
	// }
	// fmt.Println("creating user table successfully")

	
	// err = CreatePaymentTable(db)
	// if err != nil {
	// 	log.Fatal("Error creating payment table:", err)
		
	// }
	
	// fmt.Println("creating payment table successfully")

	// var payments []Payment
	// result := db.Find(&payments)
	// if result.Error != nil {
	// 	log.Fatal("Error fetching payments:", result.Error)
	// }
	// fmt.Println("no results found")

	// func CreatePaymentTable(db *gorm.DB) error {
	// 	err := db.AutoMigrate(models.Payment{})
	// 	if err != nil {
	// 		return err
	// 	}
	
	// 	return nil
	// }
	
	
	// func CreateUserTable(db *gorm.DB) error {
	// 	// AutoMigrate will create the table if it doesn't exist and apply any missing changes
	// 	err := db.AutoMigrate(models.User{})
	// 	if err != nil {
	// 		return err
	// 	}
	
	// 	return nil
	// }