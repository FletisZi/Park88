package handler

import (
	"net/http"

	"github.com/fletiszi/goteste/config"
	"github.com/fletiszi/goteste/schemas"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "layout.html", gin.H{
		"title": "Página Inicial",
	})
}

func PageCreateEstacionamentos(c *gin.Context) {
	c.HTML(http.StatusOK, "pagecreateestacionamento.html", gin.H{
		"title": "Criar novo estacionamento",
	})
}

func PageUpdateEstacionamentos(c *gin.Context) {
	c.HTML(http.StatusOK, "pageupdateestacionamento.html", gin.H{
		"title": "Atualizar estacionamento",
	})
}

func CreateEstacionamentos(c *gin.Context) {
	db := config.GetDB()

	var input schemas.Estacionamentos

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "JSON inválido",
			"details": err.Error(),
		})
		return
	}
	if input.Nome == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "O campo 'nome' é obrigatório",
		})
		return
	}

	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro ao criar estacionamento",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": input,
	})
}

func DeleteEstacionamentos(c *gin.Context) {
	db := config.GetDB()

	id := c.Param("id")

	result := db.Delete(&schemas.Estacionamentos{}, "id_estacionamento = ?", id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao deletar estacionamento",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Estacionamento não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"menssage": "Estacionamento Deletado com Sucesso!",
	})
}

func UpdateEstacionamentos(c *gin.Context) {
	db := config.GetDB()

	id := c.Param("id")

	// DTO de entrada (só campos atualizáveis)
	type UpdateInput struct {
		Nome     *string `json:"nome"`
		Endereco *string `json:"endereco"`
	}

	var input UpdateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "JSON inválido",
		})
		return
	}

	if input.Nome == nil && input.Endereco == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Informe ao menos um campo para atualizar",
		})
		return
	}

	updates := map[string]interface{}{}

	if input.Nome != nil {
		updates["nome"] = *input.Nome
	}

	if input.Endereco != nil {
		updates["endereco"] = *input.Endereco
	}

	result := db.Model(&schemas.Estacionamentos{}).
		Where("id_estacionamento = ?", id).
		Updates(updates)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao atualizar estacionamento",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Estacionamento não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Estacionamento atualizado com sucesso",
	})
}

func ListEstacionamentos(c *gin.Context) {
	db := config.GetDB()

	var estacionamentos []schemas.Estacionamentos
	if err := db.Find(&estacionamentos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar estacionamentos",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": estacionamentos,
	})
}
