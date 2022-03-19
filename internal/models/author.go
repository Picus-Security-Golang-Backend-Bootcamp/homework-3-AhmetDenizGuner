package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name  string
	Books []Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	//`gorm:"foreignKey:AuthorName;references:Name"`
}

//struct constructor
func NewAuthor(authorName string) *Author {

	author := &Author{
		Name: authorName,
	}
	return author
}
