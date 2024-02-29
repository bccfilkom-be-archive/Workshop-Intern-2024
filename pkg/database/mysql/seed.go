package mysql

import (
	"log"
	"math/rand"

	"github.com/Ndraaa15/workshop-bcc/entity"
	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func generateBook(db *gorm.DB) error {
	genres := []string{"Sci-Fi", "Fantasy", "Mystery", "Romance", "History"}
	var books []*entity.Book

	for i := 1; i <= 20; i++ {
		book := &entity.Book{
			ID:          uuid.New(),
			Title:       faker.Sentence(),
			Writter:     faker.Name(),
			Year:        rand.Intn(20) + 2000,
			Genre:       genres[rand.Intn(len(genres))],
			Description: faker.Paragraph(),
			Stock:       10,
		}
		books = append(books, book)
	}

	if err := db.CreateInBatches(books, 20).Error; err != nil {
		return err
	}
	return nil
}

func generateRole(db *gorm.DB) error {
	var roles []*entity.Role

	roles = append(roles,
		&entity.Role{
			ID:   1,
			Role: "admin",
		},
		&entity.Role{
			ID:   2,
			Role: "user",
		})

	if err := db.CreateInBatches(roles, 2).Error; err != nil {
		return err
	}
	return nil
}

func SeedData(db *gorm.DB) {
	var totalBook int64
	if err := db.Model(&entity.Book{}).Count(&totalBook).Error; err != nil {
		log.Fatalf("Error while counting book: %v", err)
	}

	if totalBook == 0 {
		if err := generateBook(db); err != nil {
			log.Fatalf("Error while generating book: %v", err)

		}
	}

	var totalRole int64
	if err := db.Model(&entity.Role{}).Count(&totalRole).Error; err != nil {
		log.Fatalf("Error while counting book: %v", err)
	}

	if totalRole == 0 {
		if err := generateRole(db); err != nil {
			log.Fatalf("Error while generating book: %v", err)

		}
	}
}
