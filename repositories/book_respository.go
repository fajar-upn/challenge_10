package repositories

import (
	"challenge_10/entity"
	"fmt"

	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type Repository interface {
	CreateBook(*entity.Book) (*entity.Book, error)
	GetBooks() ([]entity.Book, error)
	GetBooksByClientID(uuid.UUID) ([]entity.Book, error)
	GetBookByID(uuid.UUID) (*entity.Book, error)
	GetBookClientIDByID(uuid.UUID, uuid.UUID) (*entity.Book, error)
	UpdateBookByID(uuid.UUID, *entity.Book) (*entity.Book, error)
	DeleteBookByID(uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateBook(book *entity.Book) (*entity.Book, error) {

	err := r.db.Create(book).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *repository) GetBooks() ([]entity.Book, error) {
	var books []entity.Book

	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *repository) GetBooksByClientID(uuid uuid.UUID) ([]entity.Book, error) {
	var books []entity.Book

	err := r.db.Where("user_id = ?", uuid).Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *repository) GetBookByID(uuid uuid.UUID) (*entity.Book, error) {

	var book entity.Book

	err := r.db.First(&book, "id = ?", uuid).Error
	if err != nil {
		fmt.Println("Repository: ", err)
		return nil, err
	}

	return &book, nil
}

func (r *repository) GetBookClientIDByID(uuid uuid.UUID, user_id uuid.UUID) (*entity.Book, error) {

	var book entity.Book

	err := r.db.Where("user_id = ?", user_id).First(&book, "id = ?", uuid).Error
	if err != nil {
		fmt.Println("Repository: ", err)
		return nil, err
	}

	return &book, nil
}

func (r *repository) UpdateBookByID(uuid uuid.UUID, book *entity.Book) (*entity.Book, error) {

	err := r.db.Where("id = ?", uuid).
		Updates(entity.Book{Book_name: book.Book_name, Author: book.Author, Updated_at: book.Updated_at}).
		Error
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *repository) DeleteBookByID(uuid uuid.UUID) error {
	var book entity.Book

	err := r.db.Where("id = $1", uuid).Delete(&book).Error
	if err != nil {
		return err
	}

	return nil
}
