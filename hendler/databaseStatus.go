package hendler

import (
	"context"
	"net/http"
	"time"

	"github.com/fletiszi/goteste/config"
	"github.com/gin-gonic/gin"
)

func DatabaseStatus(c *gin.Context) {

	db := config.GetDB()
	if db == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "db not initialized"})
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "cannot get sqlDB", "error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "db ping failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"database": gin.H{
			"status": "🟢 Online",
		},
	})

}
