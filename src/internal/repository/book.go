package repository

import (
	"log"

	"github.com/Ndraaa15/workshop-bcc/src/entity"
	"gorm.io/gorm"
)

type IBookRepository interface {
	CreateBook(bookReq *entity.Book) (*entity.Book, error)
	UpdateBook(bookReq *entity.Book, id string) (*entity.Book, error)
	DeleteBook(id string) error
	GetBookByID(id string) (*entity.Book, error)
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
	if err := br.db.Create(&bookReq).Error; err != nil {
		return nil, err
	}
	return bookReq, nil
}

func (br *BookRepository) UpdateBook(bookReq *entity.Book, id string) (*entity.Book, error) {
	tx := br.db.Begin()

	var book entity.Book
	if err := tx.Where("id = ?", id).First(&book).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	book.Title = bookReq.Title
	book.Writter = bookReq.Writter
	book.Year = bookReq.Year
	book.Genre = bookReq.Genre
	book.Description = bookReq.Description
	book.Stock = bookReq.Stock

	if err := tx.Where("id = ?", id).Save(&book).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &book, nil
}

func (br *BookRepository) DeleteBook(id string) error {
	log.Println(id)
	if err := br.db.Where("id = ?", id).Delete(&entity.Book{}).Error; err != nil {
		return err
	}
	return nil
}

func (br *BookRepository) GetBookByID(id string) (*entity.Book, error) {
	var book entity.Book
	if err := br.db.Where("id = ?", id).Find(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}
