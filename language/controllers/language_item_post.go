package controllers

import (
	"fmt"
	"net/http"

	"github.com/faithcomesbyhearing/language-api/language/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func AddLanguage(c *gin.Context) {
	fmt.Println("Add Language")

	addLanguage := models.AddLanguageBody{}

	if err := c.BindJSON(&addLanguage); err != nil {
		fmt.Println("3.. can't bind body with addLanguage")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(addLanguage)
	if err := models.DB.Create(addLanguage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insert failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "language successfully created - " + addLanguage.Name})

}
