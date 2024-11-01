package datasource

import (
	"fmt"
	"log"

	"book-store/internal/models"

	"github.com/jinzhu/gorm"
)

func SetupDB(PostgresHost string, PostgresPort int, PostgresDB string, PostgresUser string, PostgresPwd string) (*gorm.DB, error) {
	prosgreConName := fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=disable", PostgresHost, PostgresPort, PostgresDB, PostgresUser, PostgresPwd)
	fmt.Println("conname is\t\t", prosgreConName)

	db, err := gorm.Open("postgres", prosgreConName)
	if err != nil {
		log.Fatalf("Failed to connect to database!")
		return nil, err
	}

	db.AutoMigrate(&models.Book{})

	return db, nil
}
