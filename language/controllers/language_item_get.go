package controllers

import (
	"log"
	"net/http"

	"github.com/faithcomesbyhearing/language-api/language/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// GET /language/:id
func GetLanguage(c *gin.Context) { // Get model if exist
	log.Printf("GetLanguage. id: %s", c.Param("id"))
	var language models.Language
	if err := models.DB.Where("id = ?", c.Param("id")).First(&language).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": language})
}
