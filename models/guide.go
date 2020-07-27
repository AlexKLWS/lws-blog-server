package models

import (
	"time"
)

type GuideData struct {
	Model
	MaterialData
	Type        LocationType   `json:"type" xml:"type"`
	Coordinates LocationCoords `json:"coordinates" xml:"coordinates"`
	Address     string         `json:"address" xml:"address"`
	Title       string         `json:"title" xml:"title"`
	Description string         `json:"description" xml:"description"`
	ImageUrl    string         `json:"imageUrl" xml:"imageUrl"`
}

type LocationCoords struct {
	Lat float32
	Lng float32
}

func (gd GuideData) GetID() uint {
	return gd.ID
}

func (gd GuideData) GetCreatedAt() time.Time {
	return gd.CreatedAt
}

func CreateGuideDataFromJoinedRecord(r JoinedArticlePage) GuideData {
	return GuideData{
		Model: Model{
			CreatedAt: r.CreatedAt,
		},
		MaterialData: CreateMaterialDataFromJoinedRecord(r),
	}
}
