package handler

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/kristain09/API4/app/feature/user"
	"github.com/kristain09/API4/helper"
	"github.com/labstack/echo/v4"
)

type userController struct {
	service user.UseCase
}

func New(us user.UseCase) user.Handler {
	return &userController{
		service: us,
	}
}

func (uc *userController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterInput{}
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("terjadi kesalahan bind", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan input dari user", nil))
		}

		err := uc.service.Register(user.Core{HP: input.HP, Nama: input.Nama, Password: input.Password})

		if err != nil {
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, err.Error(), nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusCreated, "sukses menambahkan data", nil))
	}
}

func (uc *userController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginInput
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("terjadi kesalahan bind", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan input dari user", nil))
		}

		res, err := uc.service.Login(input.Hp, input.Password)
		if err != nil {
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, err.Error(), nil))
		}

		var result = new(LoginResponse)
		token := helper.GenerateJWT(res.HP)
		result.Nama = res.Nama
		result.HP = res.HP
		result.Token = token

		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses login, gunakan token ini pada akses api selanjutnya : ", result))
	}
}

func (uc *userController) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := helper.DecodeJWT(c.Get("user").(*jwt.Token))

		res, err := uc.service.GetUserByID(userID)
		if err != nil {
			c.Logger().Error("terjadi kesalahan user-profile")
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses menampilkan profil", res))
	}
}

func (uc *userController) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		// ini harusnya gaperlu
		userID := helper.DecodeJWT(c.Get("user").(*jwt.Token))

		// if userID != "" {
		// 	return c.JSON(helper.ReponsFormat(http.StatusUnauthorized, "unauthorized", nil))
		// }

		var input RegisterInput
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("terjadi kesalahan bind", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan input dari user", nil))
		}
		newData := user.Core{
			Nama:     input.Nama,
			HP:       input.HP,
			Password: input.Password,
		}

		err := uc.service.UpdateProfile(newData, userID)
		if err != nil {
			c.Logger().Error("terjadi kesalahan pada pemanggilan udpate Profile", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat gangguan dari server", nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses mengupdate profile", nil))
	}
}

// func (uc *UserController) GetUser() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		res, err := uc.model.GetAllUser()

// 		if err != nil {
// 			c.Logger().Error("user model error ", err.Error())
// 			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
// 		}

// 		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses menampilkan data", res))
// 	}
// }
