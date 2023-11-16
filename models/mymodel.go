package models

import (
	"gorm.io/gorm"
)
type User struct {
    gorm.Model
    Name  string `gorm:"not null"`
    Email string `gorm:"not null"`
}

type Payment struct {
    gorm.Model
    UserID   uint    `gorm:"not null"`
    Amount   float64 `gorm:"not null"`
    Currency string  `gorm:"not null"`
    Status   string  `gorm:"not null"`
}

type Transaction struct {
    gorm.Model
    PaymentID uint    `gorm:"not null"`
    Type      string  `gorm:"not null"`
    Amount    float64 `gorm:"not null"`
}

type Account struct {
    gorm.Model
    UserID        uint    `gorm:"not null"`
    AccountNumber string  `gorm:"not null"`
    Balance       float64 `gorm:"not null"`
}

