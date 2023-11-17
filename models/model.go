package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string    `gorm:"not null"`
	Email    string    `gorm:"not null;unique"`
	Password string    `gorm:"not null"`
	Payments []Payment `gorm:"foreignKey:UserID"`
	Account  []Account `gorm:"foreignKey:UserID"`
	Transaction  []Account `gorm:"foreignKey:UserID"`
}

type Payment struct {
    gorm.Model
    UserID          uint          `gorm:"not null"`
    User            User          `gorm:"foreignKey:UserID"`
    Amount          float64       `gorm:"not null"`
    Currency        string        `gorm:"not null"`
    Reason          string        `gorm:"not null"`
    Status          string        `gorm:"not null"`
    ReceiverAccount string        `gorm:"not null"`
    PayerAccount    string        `gorm:"not null"`
    Trans           []Transaction `gorm:"foreignKey:PaymentID"`
}

type Transaction struct {
	gorm.Model
	PaymentID *uint   `gorm:"default:null"`
	Payment   *Payment `gorm:"foreignKey:PaymentID"`
	UserID   uint        `gorm:"not null"`
	User     User        `gorm:"foreignKey:UserID"`
	Type      string  `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
}

type Account struct {
	gorm.Model
	UserID        uint    `gorm:"not null"`
	User          User    `gorm:"foreignKey:UserID"`
	AccountNumber string  `gorm:"not null;unique"`
	Balance       float64 `gorm:"not null"`
}


type Demo1 struct {
	gorm.Model
	Demo1name       float64 `gorm:"default:null"`
}
