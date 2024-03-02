package service

import (
	"errors"
	"time"

	"github.com/Ndraaa15/workshop-bcc/entity"
	"github.com/Ndraaa15/workshop-bcc/internal/repository"
	"github.com/Ndraaa15/workshop-bcc/model"
)

type IRentService interface {
	RentBook(param model.RentBook) error
}

type RentService struct {
	rent repository.IRentRepository
	user repository.IUserRepository
	book repository.IBookRepository
}

func NewRentService(rent repository.IRentRepository, user repository.IUserRepository, book repository.IBookRepository) IRentService {
	return &RentService{
		rent: rent,
		book: book,
		user: user,
	}
}

func (r *RentService) RentBook(param model.RentBook) error {
	_, err := r.user.GetUser(model.UserParam{
		ID: param.UserID,
	})
	if err != nil {
		return err
	}

	book, err := r.book.GetBookByID(param.BookID.String())
	if err != nil {
		return err
	}

	if param.Total > int(book.Stock) {
		return errors.New("total rent is bigger than total stock")
	}

	rent := entity.Rent{
		UserID:   param.UserID,
		BookID:   param.BookID,
		Total:    param.Total,
		ReturnAt: time.Now().AddDate(0, 0, 7),
	}

	err = r.rent.Create(rent)
	if err != nil {
		return err
	}

	_, err = r.book.UpdateBook(&model.UpdateBook{Stock: book.Stock - uint(param.Total)}, book.ID.String())
	if err != nil {
		return err
	}

	return nil
}
