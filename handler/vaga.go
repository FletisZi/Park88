package handler

import (
	"net/http"

	"github.com/fletiszi/goteste/config"
	"github.com/fletiszi/goteste/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PageVagasEstacionamento(c *gin.Context) {
	c.HTML(http.StatusOK, "vagasestacionamento.html", gin.H{})
}

func PageCreateVaga(c *gin.Context) {
	c.HTML(http.StatusOK, "pagecreatevaga.html", gin.H{})
}

func CreateVagas(c *gin.Context) {
	db := config.GetDB()

	var input schemas.Vagas

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "JSON inválido",
			"details": err.Error(),
		})
		return
	}

	// Inicia transação
	tx := db.Begin()

	// 1️⃣ Cria a vaga
	if err := tx.Create(&input).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao criar a vaga",
			"details": err.Error(),
		})
		return
	}

	// 2️⃣ Incrementa total_vagas no estacionamento
	if err := tx.
		Model(&schemas.Estacionamentos{}).
		Where("id_estacionamento = ?", input.IDEstacionamento).
		UpdateColumn("total_vagas", gorm.Expr("total_vagas + ?", 1)).
		Error; err != nil {

		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao atualizar total de vagas do estacionamento",
			"details": err.Error(),
		})
		return
	}

	// 3️⃣ Confirma transação
	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{
		"data": input,
	})
}

func GetVagasStatus(c *gin.Context) {
	id := c.Param("id")
	db := config.GetDB()

	var vagas []schemas.Vagas

	result := db.
		Where("id_estacionamento = ?", id).
		Order(`
			CASE tipo_vaga
				WHEN 'normal' THEN 1
				WHEN 'PCD' THEN 2
				WHEN 'carga/descarga' THEN 3
				WHEN 'gerencia' THEN 4
				WHEN 'suporte' THEN 5
				ELSE 6
			END
		`).
		Order("numero_vaga ASC").
		Find(&vagas)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao buscar vagas",
			"details": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"estacionamento_id": id,
		"total_vagas":       len(vagas),
		"vagas":             vagas,
	})
}

func GetVagaID(c *gin.Context) {
	id := c.Param("id_vaga")
	db := config.GetDB()

	var resultado schemas.VagaDetalhada

	// Iniciamos a query a partir da tabela 'vagas'
	err := db.Table("vagas").
		Select("vagas.numero_vaga, vagas.tipo_vaga, vagas.status_vaga, veiculos.placa, veiculos.modelo, clientes.nome as nome_cliente").
		Joins("JOIN veiculos ON veiculos.placa = vagas.placa").
		Joins("JOIN clientes ON clientes.id_cliente = veiculos.id_cliente").
		Where("vagas.id_vaga = ?", id).
		Scan(&resultado).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Dados não encontrados",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": resultado,
	})
}
