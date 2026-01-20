package schemas

type Vagas struct {
	ID               int64    `gorm:"column:id_vaga;primaryKey" json:"id_vaga"`
	IDEstacionamento int64    `gorm:"column:id_estacionamento;not null" json:"id_estacionamento"`
	NumeroVaga       string   `gorm:"column:numero_vaga;size:10;not null" json:"numero_vaga"`
	TipoVaga         string   `gorm:"column:tipo_vaga;type:tipo_vaga_enum;not null" json:"tipo_vaga"`
	StatusVaga       string   `gorm:"column:status_vaga;type:status_vaga_enum;not null" json:"status_vaga"`
	Observacao       *string  `gorm:"column:observacao" json:"observacao"`
	Placa            *string  `gorm:"column:placa;size:10" json:"placa"`
	Veiculo          Veiculos `gorm:"foreignKey:Placa;references:Placa" json:"veiculo"`
}

type VagaDetalhada struct {
	NumeroVaga  string `json:"numero_vaga"`
	TipoVaga    string `json:"tipo_vaga"`
	StatusVaga  string `json:"status_vaga"`
	Placa       string `json:"placa"`
	Modelo      string `json:"modelo"`
	NomeCliente string `json:"nome_cliente"`
}
