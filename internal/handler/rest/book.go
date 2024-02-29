package rest

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Ndraaa15/workshop-bcc/model"
	"github.com/Ndraaa15/workshop-bcc/pkg/response"
	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateBook(ctx *gin.Context) {
	var bookReq model.CreateBook

	if err := ctx.ShouldBindJSON(&bookReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	book, err := r.service.BookService.CreateBook(&bookReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create book", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to create book", book)
}

func (r *Rest) GetBookByID(ctx *gin.Context) {
	bookID := ctx.Param("id")

	book, err := r.service.BookService.GetBookByID(bookID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to get book", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get book", book)
}

func (r *Rest) DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	log.Println(bookID)
	err := r.service.BookService.DeleteBook(bookID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete book", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete book", nil)
}

func (r *Rest) UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("id")

	var bookReq model.UpdateBook
	if err := ctx.ShouldBindJSON(&bookReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
		return
	}

	book, err := r.service.BookService.UpdateBook(&bookReq, bookID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update book", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to update book", book)
}

func (r *Rest) GetAllBook(ctx *gin.Context) {
	pageQuery := ctx.Query("page")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Failed to bind request", err)
	}

	book, err := r.service.BookService.GetAllBook(page)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get all book", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to get all book", book)
}
