package repository

import (
	"log"

	"github.com/kristain09/API4/app/feature/book"
	"github.com/kristain09/API4/app/feature/book/handler"
	"gorm.io/gorm"
)

type bookModel struct {
	db *gorm.DB
}

func New(d *gorm.DB) book.Repository {
	return &bookModel{
		db: d,
	}
}

func (bm *bookModel) Insert(newBook book.Core, userID string) (book.Core, error) {
	var insertData Book
	insertData.Judul = newBook.Judul
	insertData.Penerbit = newBook.Penerbit
	insertData.Tahun = newBook.Tahun
	insertData.UserID = userID

	if err := bm.db.Table("books").Create(&insertData).Error; err != nil {
		log.Println("Terjadi error saat create Book", err.Error())
		return book.Core{}, err
	}

	return newBook, nil
}

func (bm *bookModel) GetAll() (any, error) {
	res := []handler.ExpectedRespond{}

	err := bm.db.Table("books").Select("books.id as id, books.judul as judul, books.tahun, users.nama as nama").Joins("JOIN users on users.hp = books.user_id").Scan(&res).Error

	if err != nil {
		log.Println("Terjadi error saat select Book ", err.Error())
		return nil, err
	}

	return res, nil
}

func (bm *bookModel) Update(propUpdate book.Core, user_id string) error {
	var res Book
	if err := bm.db.Where("judul = ?", propUpdate.Judul).First(&res).Error; err != nil {
		log.Println("Terjadi error saat pencarian buku")
		return err
	}
	res.Judul = propUpdate.Judul
	res.Penerbit = propUpdate.Penerbit
	res.Tahun = propUpdate.Tahun

	if err := bm.db.Save(&res).Error; err != nil {
		log.Println("Terjadi error saat pencarian buku")
		return err
	}
	return nil
}

// func (bm *bookModel) GetBookByID(id uint) (Book, error) {
// 	res := Book{}

// 	if err := bm.db.Find(&res, id).Error; err != nil {
// 		log.Println("Terjadi error saat select Book ", err.Error())
// 		return Book{}, err
// 	}

// 	return res, nil
// }

// func (bm *BookModel) GetAllBook() ([]Book, error) {

// 	res := []Book{}

// 	// Judul Buku, Tahun Terbit, NAMA YG MENGINPUT BUKU
// 	if err := bm.db.Select("judul, tahun, penerbit").Find(&res).Error; err != nil {
// 		log.Println("Terjadi error saat select Book ", err.Error())
// 		return nil, err
// 	}

// 	return res, nil
// }
