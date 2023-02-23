package models

type Video struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	View  int    `json:"view"`
}

func NewModel() *Video {
	return &Video{}
}
