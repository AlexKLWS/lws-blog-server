package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ArticleData struct {
	gorm.Model
	MaterialData
	ArticleText string `json:"articleText" xml:"articleText"`
}

func (ad ArticleData) getID() uint {
	return ad.ID
}

func (ad ArticleData) getCreatedAt() time.Time {
	return ad.CreatedAt
}