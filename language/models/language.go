package models

type AddLanguageBody struct {
	Name string `db:"name" json:"name" binding:"required"`
	Code string `db:"code" json:"code" binding:"required"`
}

func (AddLanguageBody) TableName() string {
	return "fcbhLanguage"
}

type UpdateLanguage struct {
	Name             string `db:"name" json:"name"`
	Code             string `db:"code" json:"code"`
	Iso3             string `db:"iso3" json:"iso3"`
	Rolv_id          string `db:"rolv_id" json:"rolv_id"`
	Glotto_id        string `db:"glotto_id" json:"glotto_id"`
	Speakers         string `db:"speakers" json:"speakers"`
	LanguageDirector string `db:"languageDirector" json:"languageDirector"`
}

func (UpdateLanguage) TableName() string {
	return "fcbhLanguage"
}

type Language struct {
	Id               int64  `db:"Id" json:"id" gorm:"<-:create"`
	Name             string `db:"name" json:"name"`
	Code             string `db:"code" json:"code"`
	Iso3             string `db:"iso3" json:"iso3"`
	Rolv_id          string `db:"rolv_id" json:"rolv_id"`
	Glotto_id        string `db:"glotto_id" json:"glotto_id"`
	Speakers         string `db:"speakers" json:"speakers"`
	LanguageDirector string `db:"languageDirector" json:"languageDirector"`
}

func (Language) TableName() string {
	return "fcbhLanguage"
}
