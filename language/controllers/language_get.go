package controllers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/faithcomesbyhearing/language-api/language/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func FindLanguages(c *gin.Context) {
	log.Println("FindLanguages")

	var db *gorm.DB = models.DB

	params, _ := url.ParseQuery(c.Request.URL.RawQuery)

	tx := db.Session(&gorm.Session{})
	for key, value := range params {
		if key == "iso" && len(value) > 0 {
			tx = tx.Where("iso3 = ?", value)
		} else if key == "name" && len(value) > 0 {
			tx = tx.Where("name = ?", value)
		}
	}

	tx = tx.Where("category = ?", "primary")

	type Result struct {
		Id          int64  `json:"id"`
		Name        string `json:"name"`
		Code        string `json:"code"`
		Iso3        string `json:"iso3"`
		CountryName string `json:"country" gorm:"column:LanguageCountry__countryName"`
	}
	var results []Result

	tx.Model(&models.Language{}).Joins("LanguageCountry").Find(&results)
	c.JSON(http.StatusOK, gin.H{"data": &results})
}
