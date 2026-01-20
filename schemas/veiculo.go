package schemas

type Veiculos struct {
	Placa     string   `gorm:"column:placa;size:10;not null;unique" json:"placa"`
	IDCliente int      `gorm:"column:id_cliente;not null" json:"id_cliente"`
	Modelo    string   `gorm:"column:modelo;size:50;not null" json:"modelo"`
	Cor       string   `gorm:"column:cor;size:30;not null" json:"cor"`
	Cliente   Clientes `gorm:"foreignKey:IDCliente" json:"cliente"`
}
