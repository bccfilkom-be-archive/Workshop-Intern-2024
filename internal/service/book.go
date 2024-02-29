package service

import (
	"github.com/Ndraaa15/workshop-bcc/entity"
	"github.com/Ndraaa15/workshop-bcc/internal/repository"
	"github.com/Ndraaa15/workshop-bcc/model"
	"github.com/google/uuid"
)

type IBookService interface {
	CreateBook(bookReq *model.CreateBook) (*entity.Book, error)
	GetBookByID(id string) (*entity.Book, error)
	DeleteBook(id string) error
	UpdateBook(bookReq *model.UpdateBook, id string) (*entity.Book, error)
	GetAllBook(page int) ([]*entity.Book, error)
}

type BookService struct {
	br repository.IBookRepository
}

func NewBookService(br repository.IBookRepository) IBookService {
	return &BookService{
		br: br,
	}
}

func (bs *BookService) CreateBook(bookReq *model.CreateBook) (*entity.Book, error) {
	bookParse := &entity.Book{
		ID:          uuid.New(),
		Title:       bookReq.Title,
		Writter:     bookReq.Writter,
		Year:        bookReq.Year,
		Genre:       bookReq.Genre,
		Description: bookReq.Description,
		Stock:       bookReq.Stock,
	}

	book, err := bs.br.CreateBook(bookParse)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (bs *BookService) GetBookByID(id string) (*entity.Book, error) {
	book, err := bs.br.GetBookByID(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (bs *BookService) DeleteBook(id string) error {
	err := bs.br.DeleteBook(id)
	if err != nil {
		return err
	}
	return nil
}

func (bs *BookService) UpdateBook(bookReq *model.UpdateBook, id string) (*entity.Book, error) {
	bookParse := &entity.Book{
		Title:       bookReq.Title,
		Writter:     bookReq.Writter,
		Year:        bookReq.Year,
		Genre:       bookReq.Genre,
		Description: bookReq.Description,
		Stock:       bookReq.Stock,
	}

	book, err := bs.br.UpdateBook(bookParse, id)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (bs *BookService) GetAllBook(page int) ([]*entity.Book, error) {
	limit := 5
	offset := (page - 1) * limit

	books, err := bs.br.GetAllBook(limit, offset)
	if err != nil {
		return nil, err
	}

	return books, nil
}
