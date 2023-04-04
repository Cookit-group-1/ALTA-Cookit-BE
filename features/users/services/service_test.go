package services

import (
	"alta-cookit-be/features/users"
	"alta-cookit-be/mocks"
	"alta-cookit-be/utils/helpers"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	repo := mocks.NewUserData(t)
	password := "Alf12345"
	hash, _ := helpers.GeneratePassword(password)
	newUser := users.Core{
		Email:    "alfian12345@gmail.com",
		Username: "alfianad",
	}
	expectedData := users.Core{
		Email:    "alfian12345@gmail.com",
		Username: "alfianad",
	}

	t.Run("success register", func(t *testing.T) {
		newUser.Password = hash
		repo.On("Register", newUser).Return(expectedData, nil).Once()
		srv := New(repo)
		newUser.Password = password
		res, err := srv.Register(newUser)
		assert.Nil(t, err)
		assert.Equal(t, expectedData.ID, res.ID)
		assert.Equal(t, expectedData.Username, res.Username)
		assert.Equal(t, expectedData.Email, res.Email)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicate email", func(t *testing.T) {
		newUser.Password = hash
		repo.On("Register", newUser).Return(users.Core{}, errors.New("Duplicate users.email")).Once()
		srv := New(repo)
		newUser.Password = password
		res, err := srv.Register(newUser)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "email already exist")
		assert.Equal(t, res.Email, "")
		repo.AssertExpectations(t)

	})

	t.Run("server problem", func(t *testing.T) {
		newUser.Password = hash
		repo.On("Register", newUser).Return(users.Core{}, errors.New("server error")).Once()
		srv := New(repo)
		newUser.Password = password
		res, err := srv.Register(newUser)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res.Username, "")
		repo.AssertExpectations(t)

	})

	t.Run("validation problem", func(t *testing.T) {
		newUser.Password = hash
		srv := New(repo)
		newUser.Password = password
		res, err := srv.Register(newUser)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "validation")
		assert.Equal(t, res.Username, "")

	})

}

func TestLogin(t *testing.T) {
	data := mocks.NewUserData(t)
	// input dan respond untuk mock data
	password := "Amr12345"
	hashed, _ := helpers.GeneratePassword(password)
	inputData := users.Core{
		Email:    "jerr@alterra.id",
		Password: password,
	}
	expectedData := users.Core{
		Email:    "alfian12345@gmail.com",
		Username: "alfianad",
		Password: hashed,
	}
	t.Run("succcess login", func(t *testing.T) {

		// res dari data akan mengembalik password yang sudah di hash
		data.On("Login", inputData.Username).Return(expectedData, nil).Once()
		srv := New(data)
		inputData.Password = password
		token, res, err := srv.Login(inputData.Username, inputData.Password)
		assert.Nil(t, err)
		assert.Equal(t, expectedData.Username, res.Username)
		assert.NotNil(t, token)
		data.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		data.On("Login", inputData.Username).Return(users.Core{}, errors.New("server error")).Once()
		srv := New(data)
		token, res, err := srv.Login(inputData.Username, inputData.Password)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		data.On("Login", inputData.Username).Return(users.Core{}, errors.New("data not found")).Once()
		srv := New(data)
		token, res, err := srv.Login(inputData.Username, inputData.Password)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("wrong password", func(t *testing.T) {
		data.On("Login", inputData.Username).Return(expectedData, nil)
		srv := New(data)
		token, res, err := srv.Login(inputData.Username, "Abe12345")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "wrong password")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

}
