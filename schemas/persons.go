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

type Clientes struct {
	ID       int    `gorm:"column:id_cliente;primaryKey" json:"id_cliente"`
	Nome     string `gorm:"column:nome;size:100;not null" json:"nome"`
	CPF      string `gorm:"column:documento;size:20;not null;unique" json:"documento"`
	Telefone string `gorm:"column:telefone;size:20" json:"telefone"`
}
