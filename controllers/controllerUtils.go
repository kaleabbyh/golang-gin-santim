package controllers

import (
	"log"
	"time"

	"github.com/google/uuid"
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
type Payment struct {
	ID        			uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt 			time.Time
	UpdatedAt 			time.Time
	DeletedAt 			gorm.DeletedAt `gorm:"index"`
	UserID          	uuid.UUID      `gorm:"type:uuid;not null"`
	Amount          	float64		   `gorm:"not null"`
	Currency        	string         `gorm:"not null"`
	Reason          	string         `gorm:"not null"`
	Status          	string         `gorm:"not null"`
	ReceiverAccount 	string    	   `gorm:"not null"`
	PayerAccount    	string  	   `gorm:"not null"`
}
type CustomTime time.Time

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(ct).Format("\"2006-01-02T15:04:05Z07:00\"")), nil
}
type Transaction struct {
	ID        			uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt 			time.Time
	UpdatedAt 			time.Time
	DeletedAt 			gorm.DeletedAt `gorm:"index"`
	PaymentID 			uuid.UUID      `gorm:"type:uuid;not null"`
	UserID    			uuid.UUID      `gorm:"type:uuid;not null"`
	Type      			string  	   `gorm:"not null"`
	Amount    			float64 	   `gorm:"not null"`
	TranferedFrom   	string  	   `gorm:"default:null"`
	TranferedTo  		string   	   `gorm:"default:null"`
}
type PaymentResponse struct {
	Payment 			Payment 	 `json:"payment"`
	User   				User    	 `json:"user"`
	Message 			string  	 `json:"message"`
}


//account controllers
type User struct {
	ID        		uuid.UUID     	 `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt   `gorm:"index"`
	Name     		string   		 `gorm:"not null"`
	Email    		string   		 `gorm:"not null"`
	Password 		string   		 `gorm:"not null"`
	Role     		string   		 `gorm:"not null"`
}

type Account struct {
	ID       		uuid.UUID      	`gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt 	`gorm:"index"`
	UserID     		uuid.UUID  		`gorm:"type:uuid;not null"`
	AccountNumber	string  	 	`gorm:"not null"`
	Balance      	float64 		`gorm:"not null"`
	CreatedBy    	uuid.UUID  		`gorm:"type:uuid;not null"`
}

type AccountResponse struct {
	Account  		Account 		`json:"account"`
	User     		User    		`json:"user"`
	Message  		string  		`json:"message"`
}

type AccountResponses struct {
	Account  		[]Account		`json:"account"`
	User     		User   	 		`json:"user"`
	Message  		string 	 		`json:"message"`
}

