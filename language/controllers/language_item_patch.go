package controllers

import (
	"fmt"
	"net/http"

	"github.com/faithcomesbyhearing/language-api/language/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func UpdateLanguage(c *gin.Context) {
	fmt.Println("Update Language")
	id := c.Param("id")

	language := models.Language{}

	if err := c.BindJSON(&language); err != nil {
		fmt.Println("3.. can't bind body for updateLanguage")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(language)
	if err := models.DB.Where("id = ?", id).UpdateColumns(language).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "language successfully updated  "})

}
