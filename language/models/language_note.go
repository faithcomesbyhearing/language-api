package models

import "time"

/*
A note on GORM naming conventions:
By default, GORM uses ID as primary key, pluralized struct name to snake_cases as table name, snake_case as column name,
and uses CreatedAt, UpdatedAt to track creating/updating time
*/

type AddLanguageNote struct {
	LanguageId int    `gorm:"column:languageId" binding:"required"`
	Author     string `binding:"required"`
	Note       string `binding:"required"`
}

func (AddLanguageNote) TableName() string {
	return "languageNote"
}

type LanguageNote struct {
	Id         int       `gorm:"<-:create"`
	LanguageId int       `gorm:"column:languageId,<-:create" json:"language_id"`
	EntryDate  time.Time `gorm:"<-:create"`
	Author     string    `gorm:"<-:create"`
	Note       string
}

func (LanguageNote) TableName() string {
	return "languageNote"
}
