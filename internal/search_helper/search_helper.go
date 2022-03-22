package search_helper

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/models"
	model "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/models"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/repositories/author_repository"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/repositories/book_repository"
)

//this method search the matches between the given string and bookList element, and return the matches list
func Search(searchItems []string, authorRepository author_repository.AuthorRepository, bookRepository book_repository.BookRepository) (map[string][]models.Book, error) {

	//creating search string using by program arguments
	//BuildSearchItem method get arguments and return the string these arguments
	searchItem := buildSearchItem(searchItems)

	//check string that will be searched is equal or bigger than  3 char
	if len(searchItem) < 3 {
		err := errors.New("Lütfen daha uzun bir keime giriniz!")
		return nil, err
	}

	//search key item in author names
	authorSlice, err := authorRepository.FindAllByKey(searchItem)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	////search key item in books information
	bookSlice, err := bookRepository.FindAllByKey(searchItem)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	//to compresion and merging two results of key in book and author info
	var resultMap = map[string][]models.Book{}
	bookIDs := []int{}

	//add book info to result map
	for _, author := range authorSlice {
		for _, book := range author.Books {
			if _, ok := resultMap[author.Name]; ok {
				resultMap[author.Name] = append(resultMap[author.Name], book)
				bookIDs = append(bookIDs, int(book.ID))
			} else {
				resultMap[author.Name] = make([]model.Book, 0)
				resultMap[author.Name] = append(resultMap[author.Name], book)
				bookIDs = append(bookIDs, int(book.ID))
			}
		}
	}

	//compare is there any duplicate book , if there is not, add result map
	for _, book := range bookSlice {
		authorName := authorRepository.FindAuthorNameByID(int(book.AuthorID))
		if _, ok := resultMap[authorName]; ok {
			if !contains(bookIDs, int(book.ID)) {
				resultMap[authorName] = append(resultMap[authorName], book)
				bookIDs = append(bookIDs, int(book.ID))
			}
		} else {
			resultMap[authorName] = make([]model.Book, 0)
			resultMap[authorName] = append(resultMap[authorName], book)
			bookIDs = append(bookIDs, int(book.ID))
		}
	}

	//set error, if there is no result
	if len(resultMap) == 0 {
		err := errors.New("Aradiginiz kitap bulunamadi lutfen baska bir kelime/kelimelerle deneyiniz!")
		return nil, err
	}

	return resultMap, nil

}

//this method return a string using by elements of given string list
func buildSearchItem(argumentSlice []string) string {
	return strings.TrimSpace(strings.Join(argumentSlice, " "))
}

func contains(list []int, i int) bool {
	for _, v := range list {
		if i == v {
			return true
		}
	}
	return false
}
