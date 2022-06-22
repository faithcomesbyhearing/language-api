package controllers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/faithcomesbyhearing/language-api/language/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// GET /language
// Get all languages, filtered
func FindLanguages(c *gin.Context) {
	log.Println("FindLanguages")
	params, _ := url.ParseQuery(c.Request.URL.RawQuery)
	fmt.Println(params)
	fmt.Println(params["iso"][0])
	var languages []models.Language
	var db *gorm.DB = models.DB

	for key, value := range params {
		if key == "iso" && len(value) > 0 {
			db = db.Where("iso3 = ?", value)
		} else if key == "name" && len(value) > 0 {
			db = db.Where("name = ?", value)
		}
	}
	db.Find(&languages)

	c.JSON(http.StatusOK, gin.H{"data": languages})
}
