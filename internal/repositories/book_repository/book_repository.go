package book_repository

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) FindAll() ([]models.Book, error) {
	var books []models.Book

	result := r.db.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (r *BookRepository) InsertInitialData(books []models.Book) {

	for _, book := range books {
		r.db.Create(&book)
	}
}

func (r *BookRepository) FindAllByKey(key string) ([]models.Book, error) {
	var books []models.Book
	result := r.db.Table("table_book").Where("Name ILIKE ?", "%"+key+"%").Or("isb_nnum ILIKE ?", "%"+key+"%").Or("stock_code ILIKE ?", "%"+key+"%").Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (r *BookRepository) DeleteById(id int) error {
	result := r.db.Delete(&models.Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *BookRepository) FindById(id int) (models.Book, error) {
	var book models.Book
	result := r.db.First(&book, id)
	if result.Error != nil {
		return models.Book{}, result.Error
	}
	return book, nil
}

func (r *BookRepository) Update(b models.Book) error {
	result := r.db.Save(b)
	//r.db.Model(&c).Update("name", "deneme")

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *BookRepository) Create(b models.Book) error {
	result := r.db.Create(b)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *BookRepository) FindByName(name string) (models.Book, error) {
	var book models.Book
	result := r.db.Where("Name = ?", name).Find(&book)
	if result.Error != nil {
		return models.Book{}, result.Error
	}
	return book, nil
}
