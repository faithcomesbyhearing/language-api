package controllers

import (
	_ "github.com/go-sql-driver/mysql"
)

// func GetLanguageDetail(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var Language models.Language
// 	err := dbmap.SelectOne(&Language, "SELECT * FROM LANGUAGE.fcbhLanguage WHERE id=? LIMIT 1", id)

// 	if err == nil {
// 		Language_id, _ := strconv.ParseInt(id, 0, 64)

// 		content := &models.Language{
// 			Id:      Language_id,
// 			Name:    Language.Name,
// 			Code:    Language.Code,
// 			Iso3:    Language.Iso3,
// 			Rolv_id: Language.Rolv_id,
// 		}
// 		c.JSON(200, content)
// 	} else {
// 		c.JSON(404, gin.H{"error": "Language not found"})
// 	}
// }
// func GetLanguage(c *gin.Context) {

// 	var Language []models.Language
// 	_, err := dbmap.Select(&Language, "select * from LANGUAGE.fcbhLanguage")

// 	if err == nil {
// 		c.JSON(200, Language)
// 	} else {
// 		log.Fatal(err)
// 		c.JSON(404, gin.H{"error": "Language not found"})
// 	}
// }

// func PostLanguage(c *gin.Context) {
// 	var Language models.Language
// 	c.Bind(&Language)

// 	log.Println(Language)

// 	if Language.Name != "" && Language.Code != "" {

// 		if insert, _ := dbmap.Exec(`INSERT INTO LANGUAGE.fcbhLanguage (name, code) VALUES (?, ?)`, Language.Name, Language.Code); insert != nil {
// 			Language_id, err := insert.LastInsertId()
// 			if err == nil {
// 				content := &models.Language{
// 					Id:   Language_id,
// 					Name: Language.Name,
// 					Code: Language.Code,
// 				}
// 				c.JSON(201, content)
// 			} else {
// 				checkErr(err, "Insert failed")
// 			}
// 		}

// 	} else {
// 		c.JSON(400, gin.H{"error": "Fields are empty"})
// 	}

// }

// func UpdateLanguage(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var Language models.Language
// 	err := dbmap.SelectOne(&Language, "SELECT * FROM LANGUAGE.fcbhLanguage WHERE id=?", id)

// 	if err == nil {
// 		var input models.Language
// 		c.Bind(&input)

// 		Language_id, _ := strconv.ParseInt(id, 0, 64)

// 		Language := models.Language{
// 			Id:               Language_id,
// 			Name:             input.Name,
// 			Code:             input.Code,
// 			Iso3:             input.Iso3,
// 			Rolv_id:          input.Rolv_id,
// 			Glotto_id:        input.Glotto_id,
// 			Speakers:         input.Speakers,
// 			LanguageDirector: input.LanguageDirector,
// 		}

// 		if Language.Iso3 != "" {
// 			_, err = dbmap.Update(&Language)

// 			if err == nil {
// 				c.JSON(200, Language)
// 			} else {
// 				checkErr(err, "Updated failed")
// 			}

// 		} else {
// 			log.Fatal(err)
// 			c.JSON(400, gin.H{"error": "fields are empty"})
// 		}

// 	} else {
// 		log.Fatal(err)
// 		c.JSON(404, gin.H{"error": "Language not found"})
// 	}
// }
