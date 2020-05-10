package models

type MaterialData struct {
	Name        string   `json:"name" xml:"name"`
	Subtitle    string   `json:"subtitle" xml:"subtitle"`
	Category    Category `json:"category" xml:"category"`
	Icon        IconData `json:"icon" xml:"icon" gorm:"foreignkey:IconRefer"`
	ReferenceId string   `json:"referenceId" xml:"referenceId" gorm:"unique;not null"`
	IconRefer   uint     `json:"-"`
}
