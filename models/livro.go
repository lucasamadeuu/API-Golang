package models

type Livro struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}
