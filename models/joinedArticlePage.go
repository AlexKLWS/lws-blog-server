package models

import "time"

type JoinedArticlePage struct {
	CreatedAt   time.Time
	Name        string
	Subtitle    string
	Category    Category
	PageURL     string
	Data        string
	Height      string
	Width       string
	ReferenceId string
}
