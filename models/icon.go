package models

type IconData struct {
	Model
	Data   string `json:"data" xml:"data"`
	Height string `json:"height" xml:"height"`
	Width  string `json:"width" xml:"width"`
}
