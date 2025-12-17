package schemas

type Estacionamentos struct {
	ID       int64   `gorm:"column:id_estacionamento;primaryKey;autoIncrement" json:"id_estacionamento"`
	Nome     string  `gorm:"column:nome;size:100;not null;unique" json:"nome"`
	Endereco *string `gorm:"column:endereco;size:255" json:"endereco"`
}
