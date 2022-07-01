package models

//struct com as informações do livro
type Livro struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}
