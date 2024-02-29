package mysql

import (
	"math/rand"
	"strconv"

	"github.com/Ndraaa15/workshop-bcc/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Function to generate dummy book data
func GenerateBook(db *gorm.DB) error {
	genres := []string{"Sci-Fi", "Fantasy", "Mystery", "Romance", "History"}
	var books []*entity.Book

	for i := 1; i <= 20; i++ {
		book := &entity.Book{
			ID:          uuid.New(),
			Title:       "Book " + strconv.Itoa(i),
			Writter:     "Writter " + strconv.Itoa(i),
			Year:        2021,
			Genre:       genres[rand.Intn(len(genres))],
			Description: "Description " + strconv.Itoa(i),
			Stock:       10,
		}
		books = append(books, book)
	}

	if err := db.CreateInBatches(books, 20).Error; err != nil {
		return err
	}
	return nil
}
