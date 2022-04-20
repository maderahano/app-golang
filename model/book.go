package model

type Book struct {
	ID       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Author   string `json:"author" form:"author"`
	Category string `json:"category" form:"category"`
}
