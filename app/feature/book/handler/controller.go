package handler

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/kristain09/API4/app/feature/book"
	"github.com/kristain09/API4/helper"
	"github.com/labstack/echo/v4"
)

type bookController struct {
	srv book.UseCase
}

func New(s book.UseCase) book.Handler {
	return &bookController{
		srv: s,
	}
}

func (bc *bookController) AddBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := helper.DecodeJWT(c.Get("user").(*jwt.Token))

		input := BookRequest{}
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("terjadi kesalahan bind", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan input dari Book", nil))
		}

		res, err := bc.srv.AddBook(book.Core{Penerbit: input.Penerbit, Judul: input.Judul, Tahun: input.Tahun}, userID)

		if err != nil {
			c.Logger().Error("terjadi kesalahan", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusCreated, "sukses menambahkan data", res))
	}
}

func (bc *bookController) GetAllBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := bc.srv.GetAllBook()

		if err != nil {
			c.Logger().Error("Book model error ", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses menampilkan data", res))
	}
}

func (bc *bookController) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := BookRequest{}
		if err := c.Bind(&input); err != nil {
			c.Logger().Error(helper.ReponsFormat(http.StatusBadRequest, "terjadi kesalahan bind", nil))
		}
		var propUpdate book.Core
		propUpdate.Judul = input.Judul
		propUpdate.Penerbit = input.Penerbit

		user_id := helper.DecodeJWT(c.Get("book").(*jwt.Token))

		err := bc.srv.Update(propUpdate, user_id)
		if err != nil {
			c.Logger().Error("book model update error", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "Terdapat ganggung error", nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses melakukan update buku anda", nil))
	}
}

// func (bc *BookController) GetBookByID() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		inputParameter := c.Param("bookId")
// 		cnv, err := strconv.Atoi(inputParameter)
// 		if err != nil {
// 			c.Logger().Error("Input error ", err.Error())
// 			return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan pada input ID", nil))
// 		}
// 		res, err := bc.s.GetBookByID(uint(cnv))

// 		if err != nil {
// 			c.Logger().Error("Book model error ", err.Error())
// 			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
// 		}

// 		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses menampilkan data", res))
// 	}
// }
