package usecase_test

import (
	"errors"
	"testing"

	"github.com/kristain09/API4/app/feature/user"
	"github.com/kristain09/API4/app/feature/user/mocks"
	"github.com/kristain09/API4/app/feature/user/repository"
	"github.com/kristain09/API4/app/feature/user/usecase"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	//Setup
	repo := mocks.NewRepository(t)
	//connect usecase to repo mock
	ul := usecase.New(repo)
	succesCaseData := user.Core{
		Nama:     "Kristain Putra",
		HP:       "081223536464",
		Password: "7988das",
	}

	t.Run("Sukses login", func(t *testing.T) {
		// use mockery to return succes result
		repo.On("Login", succesCaseData.HP, succesCaseData.Password).Return(user.Core{Nama: "Kristain", HP: "081223536464"}, nil).Once()

		// call function //err not userd as iferr
		result, err := ul.Login("081223536464", "7988das")

		//assert (test return)
		assert.Nil(t, err)
		assert.Equal(t, "081223536464", result.HP)
		assert.Equal(t, "Kristain", result.Nama)

		//verifikasi expectations of mock repo
		repo.AssertExpectations(t)
	})

	t.Run("Password salah", func(t *testing.T) {
		// use mockerty to return wrong password
		repo.On("Login", "081223536464", "7988das7988").Return(user.Core{}, errors.New("password salah")).Once()

		// call function lagi
		result, err := ul.Login("081223536464", "7988das7988")

		assert.Error(t, err)
		assert.ErrorContains(t, err, "salah")
		assert.Empty(t, result.Nama)
		repo.AssertExpectations(t)
	})

	t.Run("Data tidak ditemukan", func(t *testing.T) {

		repo.On("Login", "12345", "7988dasdasdas").Return(user.Core{}, errors.New("data tidak ditemukan")).Once()

		result, err := ul.Login("12345", "7988dasdasdas")

		assert.Error(t, err)
		assert.ErrorContains(t, err, "data tidak ditemukan")
		assert.Empty(t, result.Nama)

		repo.AssertExpectations(t)
	})

	t.Run("Kesalahan pada sever", func(t *testing.T) {
		repo.On("Login", succesCaseData.HP, succesCaseData.Password).Return(user.Core{}, errors.New("column not exist")).Once()

		result, err := ul.Login(succesCaseData.HP, succesCaseData.Password)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Empty(t, result.Nama)

		repo.AssertExpectations(t)
	})
}

func TestRegister(t *testing.T) {
	repo := mocks.NewRepository(t)
	ul := usecase.New(repo)

	succesRegisterData := user.Core{
		Nama:     "Noval",
		HP:       "12345678910",
		Password: "noval123",
	}
	t.Run("Succes Register", func(t *testing.T) {
		repo.On("Insert", succesRegisterData).Return(succesRegisterData, nil).Once()

		err := ul.Register(succesRegisterData)

		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Register", func(t *testing.T) {
		repo.On("Insert", succesRegisterData).Return(user.Core{}, errors.New("too many values")).Once()

		err := ul.Register(succesRegisterData)

		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	//PUNYA KRISTAIN
	// t.Run("Succes Register", func(t *testing.T) {
	// 	repo.On("Insert", mock.AnythingOfTypeArgument("user.Core")).Return(nil).Once()

	// 	err := ul.Register(succesRegisterData)

	// 	assert.Nil(t, err)
	// })

	// t.Run("Repo Error", func(t *testing.T) {
	// 	repo.On("Insert", mock.AnythingOfTypeArgument("user.Core")).Return(errors.New("repository error"))

	// 	err := ul.Register(succesRegisterData)

	// 	assert.NotNil(t, err)
	// 	assert.Error(t, err)
	// 	assert.ErrorContains(t, err, "server")

	// 	repo.AssertExpectations(t)
	// })

}

func TestUpdateProfile(t *testing.T) {
	repo := mocks.NewRepository(t)
	ul := usecase.New(repo)
	newData := user.Core{
		Nama:     "Kristain",
		HP:       "12345",
		Password: "7988das",
	}
	userID := "1"

	t.Run("Gagal Register", func(t *testing.T) {
		repo.On("UpdateProfile", newData, userID).Return(nil).Once()

		err := ul.UpdateProfile(newData, userID)

		assert.Nil(t, err)
	})
	t.Run("Gagal Register", func(t *testing.T) {
		repo.On("UpdateProfile", newData, userID).Return(errors.New("terjadi kesalahan pada server")).Once()

		err := ul.UpdateProfile(newData, userID)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})
}

// func TestGetUserByID(t *testing.T){
// 	repo := mocks.NewRepository(t)
// 	ul := usecase.New(repo)
// 	c4c8adef-95e9-4e77-83c5-4e1204199ac6
// }

func TestGetUserByID(t *testing.T) {
	repo := mocks.NewRepository(t)
	ul := usecase.New(repo)
	var dbUser repository.User
	dbUser.HP = "12345"
	dbUser.Nama = "Kristain"
	dbUser.Password = "7988das"
	dbUser.ID = "c4c8adef-95e9-4e77-83c5-4e1204199ac6"
	userId := "c4c8adef-95e9-4e77-83c5-4e1204199ac6"

	var res user.Core
	res.Nama = dbUser.Nama
	res.HP = dbUser.HP
	res.Password = "7988das"

	t.Run("Sukses Cari User", func(t *testing.T) {
		repo.On("GetUserByID", userId).Return(res, nil).Once()

		result, err := ul.GetUserByID(userId)

		if assert.NotEmpty(t, result) {
			assert.Equal(t, res, result)
		}
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Gagal Cari User", func(t *testing.T) {
		repo.On("GetUserByID", userId).Return(user.Core{}, errors.New("terjadi permasalahan server")).Once()

		result, err := ul.GetUserByID(userId)

		assert.Empty(t, result)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})
}
