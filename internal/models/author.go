package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name string
}

//NewAuthor is struct constructor of Author model
func NewAuthor(authorName string) *Author {
	author := &Author{
		Name: authorName,
	}
	return author
}
