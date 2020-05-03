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