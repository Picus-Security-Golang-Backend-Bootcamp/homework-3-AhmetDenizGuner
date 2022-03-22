package author_repository

import (
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-AhmetDenizGuner/internal/models"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (r *AuthorRepository) InsertInitialData(books []models.Author) {

	for _, book := range books {
		r.db.Create(&book)
	}
}

func (r *AuthorRepository) FindAllByKey(key string) ([]models.Author, error) {
	var authors []models.Author

	result := r.db.Preload("Books").Where("Name ILIKE ?", "%"+key+"%").Find(&authors)

	if result.Error != nil {
		return nil, result.Error
	}

	return authors, nil
}

//get author name by author id
func (r *AuthorRepository) FindAuthorNameByID(id int) string {
	var author models.Author
	r.db.Where("ID = ?", id).Find(&author)
	return author.Name
}

func (r *AuthorRepository) FindByName(name string) (models.Author, error) {
	var author models.Author
	result := r.db.Where("Name = ?", name).Find(&author)
	if result.Error != nil {
		return models.Author{}, result.Error
	}
	return author, nil
}

func (r *AuthorRepository) Create(a models.Author) error {
	result := r.db.Create(a)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
