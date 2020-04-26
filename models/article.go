package models

import "github.com/jinzhu/gorm"

type ArticleData struct {
	gorm.Model
	MaterialData
	ArticleText string `json:"articleText" xml:"articleText"`
}
