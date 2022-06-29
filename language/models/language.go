package models

/*
A note on GORM naming conventions:
By default, GORM uses ID as primary key, pluralized struct name to snake_cases as table name, snake_case as column name,
and uses CreatedAt, UpdatedAt to track creating/updating time
*/

type AddLanguage struct {
	Name string `db:"name" json:"name" binding:"required"`
	Code string `db:"code" json:"code" binding:"required"`
}

func (AddLanguage) TableName() string {
	return "fcbhLanguage"
}

type UpdateLanguage struct {
	Name             string `json:"name"`
	Code             string `json:"code"`
	Iso3             string `json:"iso3"`
	Rolv_id          string `json:"rolv_id"`
	Glotto_id        string `json:"glotto_id"`
	Speakers         string `json:"speakers"`
	LanguageDirector string `gorm:"column:languageDirector" json:"languageDirector"`
}

func (UpdateLanguage) TableName() string {
	return "fcbhLanguage"
}

type Language struct {
	Id               uint            `gorm:"column:Id;primaryKey" json:"id" `
	Name             string          `json:"name"`
	Code             string          `json:"code"`
	Iso3             string          `json:"iso3"`
	Rolv_id          string          `json:"rolv_id"`
	Glotto_id        string          `json:"glotto_id"`
	Speakers         string          `json:"speakers"`
	LanguageDirector string          `gorm:"column:languageDirector" json:"language_director"`
	LanguageCountry  LanguageCountry `gorm:"foreignKey:LanguageID" json:"country"`
}

// `gorm:"foreignKey:LanguageID"`
func (Language) TableName() string {
	return "fcbhLanguage"
}

// ********************* Language Country
// CountryCategory enum
type CountryCategory int

const (
	Primary CountryCategory = iota
	Other
)

func (s CountryCategory) String() string {
	switch s {
	case Primary:
		return "primary"
	case Other:
		return "other"
	}
	return "unknown"
}

type LanguageCountry struct {
	LanguageID  uint            `gorm:"column:languageId"`
	CountryName string          `gorm:"column:countryName" json:"country_name"`
	Category    CountryCategory `gorm:"column:category" json:"category"`
}

func (LanguageCountry) TableName() string {
	return "fcbhLanguage_country"
}
