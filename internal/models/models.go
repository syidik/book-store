package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"Title"`
	Author string `json:"author"`
}

type CreateBookInput struct {
	Title  string `json:"Title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBook struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
}
