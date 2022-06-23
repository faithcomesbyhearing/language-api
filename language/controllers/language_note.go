package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/faithcomesbyhearing/language-api/language/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// GET /language/:id/note
// Get all language notes, without any filters yet
func FindLanguageNotes(c *gin.Context) {
	var languageId = c.Param("id")
	log.Println("FindLanguageNotes. language id: " + languageId)
	var languageNotes []models.LanguageNote
	var db *gorm.DB = models.DB
	db = db.Where("languageId = ?", languageId)

	db.Find(&languageNotes)

	c.JSON(http.StatusOK, gin.H{"data": languageNotes})
}

func GetLanguageNote(c *gin.Context) {
	var languageId = c.Param("id")
	var languageNoteId = c.Param("note_id")

	log.Println("FindLanguageNotes. language id: " + languageId + ", note id: " + languageNoteId)

	var languageNote models.LanguageNote

	if err := models.DB.Where("id = ?", languageNoteId).First(&languageNote).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": languageNote})
}

func AddLanguageNote(c *gin.Context) {
	var strLanguageId = c.Param("id")
	log.Println("Add Language Note, language_id: " + strLanguageId)

	languageId, err := strconv.Atoi(strLanguageId)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	addLanguageNote := models.AddLanguageNote{
		LanguageId: languageId,
		Author:     "",
		Note:       "",
	}

	if err := c.BindJSON(&addLanguageNote); err != nil {
		fmt.Println("3.. can't bind body with addLanguageNote")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var db *gorm.DB = models.DB
	fmt.Println(addLanguageNote)
	if err := db.Create(addLanguageNote).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insert failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "language successfully created by - " + addLanguageNote.Author})
}
