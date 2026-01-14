package schemas

type Estacionamentos struct {
	ID          int64   `gorm:"column:id_estacionamento;primaryKey;autoIncrement" json:"id_estacionamento"`
	Nome        string  `gorm:"column:nome;size:100;not null;unique" json:"nome"`
	Endereco    *string `gorm:"column:endereco;size:255" json:"endereco"`
	Ocupacao    int64   `gorm:"column:ocupacao;size:255" json:"ocupacao"`
	Total_Vagas int64   `gorm:"column:total_vagas;size:255" json:"total_vagas"`
}
