package models

import (
	"errors"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name          string
	StockCode     string
	ISBN          string
	PageNum       int
	StockQuantity int //Stock Quantity
	Price         float64
	AuthorID      uint
	Author        Author `gorm:"foreignKey:AuthorID"`
	//AuthorName  string `gorm:"type:varchar(100);column:AuthorName"`
}

//struct constructor
func NewBook(pageNum, stockNumber int, price float64, bookName, stockCode, isbn_num string, authorID uint) *Book {

	book := &Book{
		Name:          bookName,
		StockCode:     stockCode,
		ISBN:          isbn_num,
		PageNum:       pageNum,
		StockQuantity: stockNumber,
		Price:         price,
		AuthorID:      authorID,
	}

	return book
}

//This function check some rules and decrease the stock quantity
func (b *Book) Buy(count int) error {

	//check ıs there enough book to buy
	if count > b.StockQuantity {
		err := errors.New("Yeterli sayida kitap yoktur lutfen daha az miktarda deneyiniz!")
		return err
	}

	b.StockQuantity -= count

	return nil
}
