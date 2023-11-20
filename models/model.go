package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       		uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
	Name    		string   	   `gorm:"not null"`
	Email    		string   	   `gorm:"not null;unique"`
	Password		string    	   `gorm:"not null"`
	Role     		string         `gorm:"not null"`
	// Role 		RoleEnum `gorm:"type:role_enum"`
	Payments 		[]Payment      `gorm:"foreignKey:UserID"`
	Account  		[]Account 	   `gorm:"foreignKey:UserID"`
	Transaction 	[]Account      `gorm:"foreignKey:UserID"`
}

type Payment struct {
	ID        		uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
    UserID          uuid.UUID      `gorm:"type:uuid;not null"`
    User            User           `gorm:"foreignKey:UserID"`
    Amount          float64        `gorm:"not null"`
    Currency        string         `gorm:"not null"`
    Reason          string         `gorm:"not null"`
    Status          string         `gorm:"not null"`
    ReceiverAccount string         `gorm:"not null"`
    PayerAccount    string         `gorm:"not null"`
    Trans           []Transaction  `gorm:"foreignKey:PaymentID"`
}

type Transaction struct {
	ID       		uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt 	`gorm:"index"`
	PaymentID 		uuid.UUID       `gorm:"type:uuid;not null"`
	Payment   		*Payment 		`gorm:"foreignKey:PaymentID"`
	UserID   		uuid.UUID       `gorm:"type:uuid;not null"`
	User     		User       		`gorm:"foreignKey:UserID"`
	Type     		string 			`gorm:"not null"`
	Amount    		float64		 	`gorm:"not null"`
	TranferedFrom   string 		 	`gorm:"default:null"`
	TranferedTo   	string   		`gorm:"default:null"`
}

type Account struct {
	ID        		uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
	UserID          uuid.UUID      `gorm:"type:uuid;not null"`
	User            User   		   `gorm:"foreignKey:UserID"`
	AccountNumber   string   	   `gorm:"not null;unique"`
	Balance         float64 	   `gorm:"not null"`
	CreatedBy    	uuid.UUID      `gorm:"type:uuid;not null"`
	CreatedByUser   User           `gorm:"foreignKey:CreatedBy"`
	
}


type Demo struct {
	ID        		uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
	Demoname  		string         `gorm:"default:null"`
}


type RoleEnum string
const (
    user  RoleEnum = "user"
    admin RoleEnum = "admin"
	superadmin RoleEnum = "superadmin"
)
