package usecase

import (
	"errors"
	"strings"

	"github.com/kristain09/API4/app/feature/user"
	"github.com/labstack/gommon/log"
)

type userLogic struct {
	m user.Repository
}

func New(r user.Repository) user.UseCase {
	return &userLogic{
		m: r,
	}
}

func (ul *userLogic) Login(hp string, password string) (user.Core, error) {
	result, err := ul.m.Login(hp, password)
	if err != nil {

		if strings.Contains(err.Error(), "tidak ditemukan") {
			return user.Core{}, errors.New("data tidak ditemukan")
		} else if strings.Contains(err.Error(), "salah") {
			return user.Core{}, errors.New("password salah")
		}

		return user.Core{}, errors.New("terdapat permasalahan pada server")
	}

	return result, nil
}

func (ul *userLogic) GetUserByID(userID string) (user.Core, error) {
	result, err := ul.m.GetUserByID(userID)
	if err != nil {
		return user.Core{}, errors.New("terjadi permasalahan server")
	}

	return result, nil
}

func (ul *userLogic) Register(newUser user.Core) error {
	_, err := ul.m.Insert(newUser)
	if err != nil {
		log.Error("register logic error:", err.Error())

		return errors.New("terjadi kesalahan pada server")
	}

	return nil
}

func (ul *userLogic) UpdateProfile(newData user.Core, userID string) error {
	err := ul.m.UpdateProfile(newData, userID)
	if err != nil {
		log.Error("updateprofile logic error:", err.Error())
		return errors.New("terjadi kesalahan pada server")
	}

	return nil
}
