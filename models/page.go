package models

import (
	"time"
)

type PageData struct {
	Model
	MaterialData
	PageURL string `json:"pageURL" xml:"pageURL"`
}

func (pg PageData) GetID() uint {
	return pg.ID
}

func (pg PageData) GetCreatedAt() time.Time {
	return pg.CreatedAt
}

func CreatePageDataFromJoinedRecord(r JoinedArticlePage) PageData {
	return PageData{
		Model: Model{
			CreatedAt: r.CreatedAt,
		},
		MaterialData: CreateMaterialDataFromJoinedRecord(r),
		PageURL:      r.PageURL,
	}
}
