package database

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/models"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/repositories/author_repository"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/repositories/book_repository"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/pkg/csv_helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=root dbname=library port=5432" //sslmode=disable TimeZone=Asia/Shanghai
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "table_",
			SingularTable: true,
		},
		/*NowFunc: func() time.Time {
			return time.Now().UTC()
		},*/
	})

	if err != nil {
		panic(fmt.Sprintf("Could not connect to the database: %s", err.Error()))
	}

	DB = connection

	//connection.AutoMigrate(&models.User{})
}

func InitiliazeDatabase() error {

	dbExist := DB.Migrator().HasTable(&models.Book{})

	if !dbExist {
		DB.AutoMigrate(&models.Book{})
		DB.AutoMigrate(&models.Author{})

		authorSlice, err := csv_helper.ReadCsv("../resources/author.csv", 1)

		if err != nil {
			fmt.Println(err)
			return errors.New("CSV cannot be read!")
		}

		var authorData []models.Author
		for _, author := range authorSlice {

			newAuthor := models.NewAuthor(author[1])

			authorData = append(authorData, *newAuthor)

			//DB.Where(models.Book{Name: newAuthor.Name}).Attrs(models.Author{Name: newAuthor.Name}).FirstOrCreate(&newAuthor)

		}

		authorRepo := author_repository.NewAuthorRepository(DB)
		authorRepo.InsertInitialData(authorData)

		bookSlice, err := csv_helper.ReadCsv("../resources/book.csv", 1)

		if err != nil {
			fmt.Println(err)
			return errors.New("CSV cannot be read!")
		}

		var bookData []models.Book

		for _, book := range bookSlice {

			price, _ := strconv.ParseFloat(book[4], 64)
			pageNum, _ := strconv.Atoi(book[3])
			stockQuantity, _ := strconv.Atoi(book[5])
			authorID, _ := strconv.Atoi(book[6])
			newBook := models.NewBook(pageNum, stockQuantity, price, book[0], book[2], book[1], uint(authorID))

			bookData = append(bookData, *newBook)

			//DB.Where(models.Book{Name: newBook.Name}).Attrs(models.Book{Name: newBook.Name, StockCode: newBook.StockCode, ISBNnum: newBook.ISBNnum, PageNum: newBook.PageNum, StockNumber: newBook.StockNumber, Price: newBook.Price, AuthorID: newBook.AuthorID}).FirstOrCreate(&newBook)

		}

		bookRepo := book_repository.NewBookRepository(DB)
		bookRepo.InsertInitialData(bookData)

	}

	return nil

}
