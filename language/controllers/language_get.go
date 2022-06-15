package controllers

import (
	"log"
	"net/http"

	"github.com/faithcomesbyhearing/language-api/language/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// GET /language
// Get all languages, filtered
// FIXME: implement filter
func FindLanguages(c *gin.Context) {
	log.Print("FindLanguages")
	var languages []models.Language
	models.DB.Find(&languages)
	c.JSON(http.StatusOK, gin.H{"data": languages})
}
