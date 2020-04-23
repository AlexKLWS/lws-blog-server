package models

import "github.com/jinzhu/gorm"

type IconData struct {
	gorm.Model
	Data   string `json:"data" xml:"data"`
	Height string `json:"height" xml:"height"`
	Width  string `json:"width" xml:"width"`
}
