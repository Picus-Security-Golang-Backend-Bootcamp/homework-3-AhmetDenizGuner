package models

import (
	"errors"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string
	StockCode   string
	ISBNnum     string
	PageNum     int
	StockNumber int //Stock Quantity
	Price       float64
	AuthorID    uint
	//AuthorName  string `gorm:"type:varchar(100);column:AuthorName"`
}

//struct constructor
func NewBook(pageNum, stockNumber int, price float64, bookName, stockCode, isbn_num string, authorID uint) *Book {

	book := &Book{
		Name:        bookName,
		StockCode:   stockCode,
		ISBNnum:     isbn_num,
		PageNum:     pageNum,
		StockNumber: stockNumber,
		Price:       price,
		AuthorID:    authorID,
	}

	return book
}

//This function check some rules and decrease the stock quantity
func (b *Book) Buy(count int) error {

	//check ıs there enough book to buy
	if count > b.StockNumber {
		err := errors.New("Yeterli sayida kitap yoktur lutfen daha az miktarda deneyiniz!")
		return err
	}

	b.StockNumber -= count

	return nil
}

/*type Deletable interface {
	Delete() (string, error)
}

//This function check some rules and set the isDeleted field as true
func (b *Book) Delete() (string, error) {

	//check book is deleted before
	if b.IsDeleted {
		err := errors.New("Bu kitap zaten silinmiş tekrar silemezsiniz!")
		return "", err
	}

	b.IsDeleted = true

	return b.Name + " Basariyla silindi", nil
}*/
