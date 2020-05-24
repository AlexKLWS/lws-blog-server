package models

type PageIndex struct {
	Model
	Page  int    `json:"page" xml:"page"`
	Query string `json:"query" xml:"query"`
}
