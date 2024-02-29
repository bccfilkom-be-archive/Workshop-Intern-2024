package handler

import (
	"log"
	"net/http"

	"github.com/Ndraaa15/workshop-bcc/internal/service"
	"github.com/Ndraaa15/workshop-bcc/model"
	"github.com/Ndraaa15/workshop-bcc/pkg/response"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	BookService service.IBookService
}

func NewBookHandler(bookService service.IBookService) *BookHandler {
	return &BookHandler{
		BookService: bookService,
	}
}

func (bh *BookHandler) CreateBook(ctx *gin.Context) {
	var bookReq model.CreateBook

	if err := ctx.ShouldBindJSON(&bookReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	book, err := bh.BookService.CreateBook(&bookReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create book", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to create book", book)
}

func (bh *BookHandler) GetBookByID(ctx *gin.Context) {
	bookID := ctx.Param("id")

	book, err := bh.BookService.GetBookByID(bookID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to get book", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get book", book)
}

func (bh *BookHandler) DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	log.Println(bookID)
	err := bh.BookService.DeleteBook(bookID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete book", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete book", nil)
}

func (bh *BookHandler) UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("id")

	var bookReq model.UpdateBook
	if err := ctx.ShouldBindJSON(&bookReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	book, err := bh.BookService.UpdateBook(&bookReq, bookID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update book", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to update book", book)
}
