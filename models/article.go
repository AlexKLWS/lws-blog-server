package models

import "github.com/jinzhu/gorm"

type ArticleData struct {
	gorm.Model
	Name        string   `json:"name" xml:"name"`
	Subtitle    string   `json:"subtitle" xml:"subtitle"`
	ArticleText string   `json:"articleText" xml:"articleText"`
	Icon        IconData `json:"icon" xml:"icon" gorm:"foreignkey:IconRefer"`
	IconRefer uint
}
