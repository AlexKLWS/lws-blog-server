package models

type ArticleData struct {
	Name        string   `json:"name" xml:"name"`
	Subtitle    string   `json:"subtitle" xml:"subtitle"`
	ArticleText string   `json:"articleText" xml:"articleText"`
	Icon        IconData `json:"icon" xml:"icon"`
}
