package repository

import (
	"log"

	"github.com/Ndraaa15/workshop-bcc/entity"
	"github.com/Ndraaa15/workshop-bcc/model"
	"gorm.io/gorm"
)

type IBookRepository interface {
	CreateBook(bookReq *entity.Book) (*entity.Book, error)
	UpdateBook(bookReq *model.UpdateBook, id string) (*entity.Book, error)
	DeleteBook(id string) error
	GetBookByID(id string) (*entity.Book, error)
	GetAllBook(limit, offset int) ([]*entity.Book, error)
}

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &BookRepository{
		db: db,
	}
}

func (br *BookRepository) CreateBook(bookReq *entity.Book) (*entity.Book, error) {
	if err := br.db.Debug().Create(&bookReq).Error; err != nil {
		return nil, err
	}
	return bookReq, nil
}

func (br *BookRepository) UpdateBook(bookReq *model.UpdateBook, id string) (*entity.Book, error) {
	tx := br.db.Begin()

	var book *entity.Book
	if err := tx.Debug().Where("id = ?", id).First(&book).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	book = parseUpdateReq(book, bookReq)

	if err := tx.Debug().Where("id = ?", id).Save(&book).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return book, nil
}

func (br *BookRepository) DeleteBook(id string) error {
	log.Println(id)
	if err := br.db.Debug().Where("id = ?", id).Delete(&entity.Book{}).Error; err != nil {
		return err
	}
	return nil
}

func (br *BookRepository) GetBookByID(id string) (*entity.Book, error) {
	var book entity.Book
	if err := br.db.Debug().Where("id = ?", id).Find(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (br *BookRepository) GetAllBook(limit, offset int) ([]*entity.Book, error) {
	var books []*entity.Book
	if err := br.db.Debug().Limit(limit).Offset(offset).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func parseUpdateReq(book *entity.Book, bookReq *model.UpdateBook) *entity.Book {
	if bookReq.Title != "" {
		book.Title = bookReq.Title
	}
	if bookReq.Writter != "" {
		book.Writter = bookReq.Writter
	}
	if bookReq.Year >= 0 {
		book.Year = bookReq.Year
	}
	if bookReq.Genre != "" {
		book.Genre = bookReq.Genre
	}
	if bookReq.Description != "" {
		book.Description = bookReq.Description
	}
	if bookReq.Stock >= 0 {
		book.Stock = bookReq.Stock
	}

	return book
}
