package models

type Language struct {
	Id               int64  `db:"Id" json:"id"`
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
