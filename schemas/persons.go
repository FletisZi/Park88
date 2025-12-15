package schemas

import (
	"gorm.io/gorm"
)

type Opning struct {
	gorm.Model
	Name     string
	User     string
	Password string
	Age      int16
	Country  string
}

type User struct {
	ID    int
	Name  string
	Email string
}
