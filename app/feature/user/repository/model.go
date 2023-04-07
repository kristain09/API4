package repository

import (
	"time"

	"github.com/kristain09/API4/app/feature/book/repository"
	"gorm.io/gorm"
)

type User struct {
	ID        string
	Nama      string
	HP        string `gorm:"primaryKey;type:varchar(13);"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt    `gorm:"index"`
	Books     []repository.Book `gorm:"foreignKey:UserID"`
}
