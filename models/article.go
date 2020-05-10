package models

import (
	"time"
)

type ArticleData struct {
	Model
	MaterialData
	ArticleText string `json:"articleText" xml:"articleText"`
}

func (ad ArticleData) GetID() uint {
	return ad.ID
}

func (ad ArticleData) GetCreatedAt() time.Time {
	return ad.CreatedAt
}

func CreateArticleDataFromJoinedRecord(r JoinedArticlePage) ArticleData {
	return ArticleData{
		Model: Model{
			CreatedAt: r.CreatedAt,
		},
		MaterialData: CreateMaterialDataFromJoinedRecord(r),
	}
}
