package handler

import (
	"net/http"

	"github.com/fletiszi/goteste/config"
	"github.com/fletiszi/goteste/schemas"
	"github.com/gin-gonic/gin"
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
	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao criar a vaga",
			"details": err.Error(),
		})
		return
	}
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
