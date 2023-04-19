package service

import (
	"challenge_10/entity"
	"challenge_10/repositories"
	"time"

	"github.com/gofrs/uuid/v5"
)

type Service interface {
	CreateBook(*entity.PayloadBook) (*entity.Book, error)
	GetBooks() ([]entity.Book, error)
	GetBooksByClientID(uuid.UUID) ([]entity.Book, error)
	GetBookByID(uuid.UUID) (*entity.Book, error)
	GetBookClientIDByID(uuid.UUID, uuid.UUID) (*entity.Book, error)
	UpdateBookByID(uuid.UUID, *entity.PayloadBook) (*entity.Book, error)
	DeleteBookByID(uuid.UUID) error
}

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) *service {
	return &service{repository}
}

func (s *service) CreateBook(payloadBook *entity.PayloadBook) (*entity.Book, error) {

	uuid, _ := uuid.NewV4()

	book := &entity.Book{
		ID:         uuid,
		User_id:    payloadBook.User_id,
		Book_name:  payloadBook.Book_name,
		Author:     payloadBook.Author,
		Created_at: time.Now().Local(),
		Updated_at: time.Now().Local(),
	}

	book, err := s.repository.CreateBook(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *service) GetBooks() ([]entity.Book, error) {

	books, err := s.repository.GetBooks()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *service) GetBooksByClientID(uuid uuid.UUID) ([]entity.Book, error) {

	books, err := s.repository.GetBooksByClientID(uuid)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *service) GetBookByID(uuid uuid.UUID) (*entity.Book, error) {

	books, err := s.repository.GetBookByID(uuid)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *service) GetBookClientIDByID(uuid, user_id uuid.UUID) (*entity.Book, error) {
	books, err := s.repository.GetBookClientIDByID(uuid, user_id)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *service) UpdateBookByID(uuid uuid.UUID, payloadBook *entity.PayloadBook) (*entity.Book, error) {

	getBookByID, _ := s.repository.GetBookByID(uuid)

	if payloadBook.Book_name == "" {
		payloadBook.Book_name = getBookByID.Book_name
	}

	if payloadBook.Author == "" {
		payloadBook.Author = getBookByID.Author
	}

	book := &entity.Book{
		ID:         getBookByID.ID,
		User_id:    getBookByID.User_id,
		Book_name:  payloadBook.Book_name,
		Author:     payloadBook.Author,
		Created_at: getBookByID.Created_at,
		Updated_at: time.Now().Local(),
	}

	updateBook, err := s.repository.UpdateBookByID(uuid, book)
	if err != nil {
		return nil, err
	}

	return updateBook, nil
}

func (s *service) DeleteBookByID(uuid uuid.UUID) error {

	err := s.repository.DeleteBookByID(uuid)
	if err != nil {
		return err
	}

	return nil
}
