package models

type Login struct {
	Password string `json:"password"`
}

type Session struct {
	Token string `json:"token" xml:"token"`
}
