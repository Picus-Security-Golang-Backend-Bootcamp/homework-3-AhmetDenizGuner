package search_helper

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/models"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/repositories/author_repository"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/repositories/book_repository"
)

//this method search the matches between the given string and bookList element, and return the matches list
func Search(searchItems []string, authorRepository author_repository.AuthorRepository, bookRepository book_repository.BookRepository) ([]models.Book, error) {

	//creating search string using by program arguments
	//BuildSearchItem method get arguments and return the string these arguments
	searchItem := buildSearchItem(searchItems)

	//check string that will be searched is equal or bigger than  3 char
	if len(searchItem) < 3 {
		err := errors.New("Lütfen daha uzun bir keime giriniz!")
		return nil, err
	}

	////search key item in books information
	bookSlice, err := bookRepository.FindAllByKey(searchItem)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	//set error, if there is no result
	if len(bookSlice) == 0 {
		err := errors.New("Aradiginiz kitap bulunamadi lutfen baska bir kelime/kelimelerle deneyiniz!")
		return nil, err
	}

	return bookSlice, nil

}

//buildSearchItem return a string using by elements of given string list
func buildSearchItem(argumentSlice []string) string {
	return strings.TrimSpace(strings.Join(argumentSlice, " "))
}
