package models

import "github.com/jinzhu/gorm"

type PageData struct {
	gorm.Model
	MaterialData
	PageURL string `json:"pageURL" xml:"pageURL"`
}
