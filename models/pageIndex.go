package models

import "time"

type PageIndex struct {
	Model
	Page      int       `json:"page" xml:"page"`
	StartDate time.Time `json:"startDate" xml:"startDate"`
	EndDate   time.Time `json:"endDate" xml:"endDate"`
	Category  Category  `json:"category" xml:"category"`
}
