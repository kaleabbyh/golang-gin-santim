package controllers

import (
	"log"

	"github.com/kaleabbyh/golang-santim/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = config.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
}

//payment controller
type user struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
}

type payment struct {
	gorm.Model
	UserID          uint    `gorm:"not null"`
	Amount          float64 `gorm:"not null"`
	Currency        string  `gorm:"not null"`
	Reason          string  `gorm:"not null"`
	Status          string  `gorm:"not null"`
	ReceiverAccount string  `gorm:"not null"`
	PayerAccount    string  `gorm:"not null"`
}

type transaction struct {
	gorm.Model
	PaymentID *uint   `gorm:"default:null"`
	UserID    uint    `gorm:"not null"`
	Type      string  `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
	TranferedFrom string  `gorm:"default:null"`
	TranferedTo string    `gorm:"default:null"`
}
type paymentResponse struct {
	Payment payment `json:"payment"`
	User    user    `json:"user"`
	Message string  `json:"message"`
}



//account controllers


type User struct {
	gorm.Model
	Name     string    `gorm:"not null"`
	Email    string    `gorm:"not null"`
	Password string    `gorm:"not null"`
}

type Account struct {
	gorm.Model
	UserID        uint    `gorm:"not null"`
	AccountNumber string  `gorm:"not null"`
	Balance       float64 `gorm:"not null"`
}

type AccountResponse struct {
	Account  Account `json:"account"`
	User     User    `json:"user"`
	Message  string  `json:"message"`
}

type AccountResponses struct {
	Account  []Account `json:"account"`
	User     User    `json:"user"`
	Message  string  `json:"message"`
}

