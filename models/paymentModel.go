package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kaleabbyh/golang-santim/config"
)

type payment struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Reason string  `json:"reason"`
	Price  float64 `json:"price"`
}

func CreateTable(db *sql.DB) error {
	// Create the SQL statement to create the table
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL
	);`

	// Execute the SQL statement
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func CreatePaymentTable(db *sql.DB) error {
	// Create the SQL statement to create the table
	createTableSQL := `CREATE TABLE IF NOT EXISTS payments (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		reason VARCHAR(255) NOT NULL,
		price  INTEGER NOT NULL
	);`

	// Execute the SQL statement
	_, err := db.Exec(createTableSQL)
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
	defer db.Close()

	err = CreateTable(db)
	if err != nil {
		log.Fatal("Error  creating user table:", err)
		
	}

	fmt.Println("creating user table successfully")

	err = CreatePaymentTable(db)
	if err != nil {
		log.Fatal("Error creating payment table:", err)
		
	}
	
	fmt.Println("creating payment table successfully")

}