package command_handler

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/database"
	model "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/models"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/repositories/author_repository"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/repositories/book_repository"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/search_helper"
)

func List(args []string) {
	if len(args) == 2 {

		bookRepo := book_repository.NewBookRepository(database.DB)
		bookSlice, err := bookRepo.GetAllBooks()

		if err != nil {
			fmt.Println(err)
			return
		}

		bookPrinter(bookSlice)
	} else {
		fmt.Println("list komutu baska arguman alamaz!")
		return
	}
}

func Search(args []string) {
	//check is there any search string
	if len(args) < 3 {
		fmt.Println("Search komutunu kullanabilmek için lütfen arama yapilacak kelime veya kelimleri de giriniz!")
		return
	}

	//Searching string in book list
	//Search method use strings contains method and seek any match in book list and return the all matches
	resultMap, err := search_helper.Search(args[2:])

	//if resultSlices is empty Serach method will be return error message and program terminated
	if err != nil {
		fmt.Println(err)
		return
	}

	keys := reflect.ValueOf(resultMap).MapKeys()

	for _, authorName := range keys {
		for _, book := range resultMap[authorName.String()] {
			fmt.Printf("BookID: %d , Name: %s,  Price: %.2f, ISBN: %s, Stock Code: %s, Stock Num: %d Author Name: %s\n", book.ID, book.Name, book.Price, book.ISBNnum, book.StockCode, book.StockNumber, authorName)
		}
	}
}

func Buy(args []string) {
	//check argument count is ok
	if len(args) != 4 {
		fmt.Println("Buy komutu icin hatali sayida arguman girdiniz, lutfen 2 tam sayi giriniz!")
		return
	}

	bookId, err1 := strconv.Atoi(args[2])
	orderCount, err2 := strconv.Atoi(args[3])

	//check arguments are integers
	if err1 != nil || err2 != nil {
		fmt.Println("Lutfen arguman degerlerini tam sayi giriniz!")
		return
	}

	bookRepo := book_repository.NewBookRepository(database.DB)
	book, err := bookRepo.GetById(bookId)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Islem yapmaya calistiginiz ID sistemde yok veya silinmis lutfen baska ID ile deneyiniz")
		return
	}

	err3 := book.Buy(orderCount)

	if err3 != nil {
		fmt.Println(err)
		return
	}

	err4 := bookRepo.Update(book)

	if err4 != nil {
		fmt.Println(err4)
		return
	}

	fmt.Printf("%s 'den %d tane basariyla alindi.", book.Name, orderCount)

}

func Delete(args []string) {
	//check argument count is ok
	if len(args) != 3 {
		fmt.Println("Delete komutu icin hatali sayida arguman girdiniz, lutfen 1 tam sayi giriniz!")
		return
	}

	bookId, err1 := strconv.Atoi(args[2])

	//check arguments are integers
	if err1 != nil {
		fmt.Println("Lutfen arguman degerini tam sayi giriniz!")
		return
	}

	//check book id added at the begginig
	bookRepo := book_repository.NewBookRepository(database.DB)
	book, err := bookRepo.GetById(bookId)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Islem yapmaya calistiginiz ID sistemde yok veya zaten silinmis lutfen baska ID ile deneyiniz")
		return
	}

	err2 := bookRepo.DeleteById(int(book.ID))

	if err != nil {
		fmt.Println(err2)
		return
	}

	fmt.Println("Kitap basariyla silindi")
}

func Update(args []string) {
	//check argument count is ok
	if len(args) != 4 {
		fmt.Println("Buy komutu icin hatali sayida arguman girdiniz, lutfen 2 tam sayi giriniz!")
		return
	}

	bookId, err1 := strconv.Atoi(args[2])
	increaseCount, err2 := strconv.Atoi(args[3])

	//check arguments are integers
	if err1 != nil || err2 != nil {
		fmt.Println("Lutfen arguman degerlerini tam sayi giriniz!")
		return
	}

	//check book id added at the begginig
	bookRepo := book_repository.NewBookRepository(database.DB)
	book, err := bookRepo.GetById(bookId)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Islem yapmaya calistiginiz ID sistemde yok veya zaten silinmis lutfen baska ID ile deneyiniz")
		return
	}

	book.StockNumber += increaseCount

	err4 := bookRepo.Update(book)

	if err4 != nil {
		fmt.Println(err4)
		return
	}

	fmt.Printf("%s kitabinin stogu %d tane arttirildi.", book.Name, increaseCount)

}

func Add(args []string) {

	//check argument count is ok
	if len(args) >= 4 {
		fmt.Println("Add icin lutfen argumanlari asagidaki gibi giriniz")
		fmt.Println("go run main.go add Kitap Adi,Yazar Adi, ISBN, Stock Code, Stock Quantity, Page Number, Price")
	}

	input := strings.TrimSpace(strings.Join(args[2:], ""))
	arguments := strings.Split(input, ",")

	if len(arguments) != 7 {
		fmt.Println("Add icin lutfen argumanlari asagidaki gibi giriniz")
		fmt.Println("go run main.go add Kitap Adi,Yazar Adi, ISBN, Stock Code, Stock Quantity, Page Number, Price")
	}

	stockQuantity, err := strconv.Atoi(arguments[4])
	pageNumber, err2 := strconv.Atoi(arguments[5])
	price, err3 := strconv.ParseFloat(arguments[6], 64)

	if err != nil || err2 != nil || err3 != nil {
		fmt.Println("Lutfen argumanlarin veri tiplerine dikkat ediniz!")
		fmt.Println("go run main.go add Kitap Adi,Yazar Adi, ISBN, Stock Code, Stock Quantity, Page Number, Price")
		return
	}

	authorRepo := author_repository.NewAuthorRepository(database.DB)
	author, err4 := authorRepo.FindByName(arguments[1])

	if err4 != nil {
		author.Name = arguments[1]
		err := authorRepo.Create(author)

		if err != nil {
			fmt.Println(err)
			return
		}
		author, _ = authorRepo.FindByName(arguments[1])
	}

	bookRepo := book_repository.NewBookRepository(database.DB)
	_, err5 := bookRepo.FindByName(arguments[0])

	if err5 != nil {
		fmt.Println("Eklemeye calistiginiz kitap zaten var!!")
		return
	}

	book := model.NewBook(pageNumber, stockQuantity, price, arguments[0], arguments[2], arguments[1], author.ID)

	err6 := bookRepo.Create(*book)

	if err6 != nil {
		fmt.Println(err6)
		return
	}

	fmt.Println("Kitap basariyla eklendi!!!")

}

//this function print the list that given
func bookPrinter(books []model.Book) {
	for index, book := range books {
		fmt.Printf("%d) BookID: , Name: %s,  Price: %.2f, ISBN: %s, Stock Code: %s, Stock Num: %d\n", index+1, book.Name, book.Price, book.ISBNnum, book.StockCode, book.StockNumber)
	}
}
