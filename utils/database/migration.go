package database

import (
	"github.com/kristain09/API4/app/feature/book/repository"
	urepo "github.com/kristain09/API4/app/feature/user/repository"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	books := repository.Book{}
	users := urepo.User{}
	db.AutoMigrate(books)
	db.AutoMigrate(users)
}
