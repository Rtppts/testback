package entity

import (
	"gorm.io/gorm"
)

type Ratta struct{
	gorm.Model
	Name string `valid:"required~Name is required"`
	ID string `gorm:"uniqueIndex" valid:"required~ID is required, matches(^[BMD]\\d{7}$)~Name is not match"`
	Phone string `valid:"required~Phone is required, stringlength(10|10)~Phone is must be 10 digit"`
	Email string `valid:"required~Email is required, email~Email is invalid"`
}