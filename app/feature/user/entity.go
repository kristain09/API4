package user

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	Nama     string
	HP       string
	Password string
}

type Handler interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
	Profile() echo.HandlerFunc
	UpdateProfile() echo.HandlerFunc
}

type UseCase interface {
	Login(hp string, password string) (Core, error)
	Register(newUser Core) error
	GetUserByID(userID string) (Core, error)
	UpdateProfile(newData Core, userID string) error
}

type Repository interface {
	Insert(newUser Core) (Core, error)
	Login(hp string, password string) (Core, error)
	GetUserByID(userID string) (Core, error)
	UpdateProfile(newData Core, userID string) error
}
