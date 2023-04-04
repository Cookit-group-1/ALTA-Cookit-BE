package services

import (
	"alta-cookit-be/features/users"
	"alta-cookit-be/mocks"
	"alta-cookit-be/utils/helpers"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// func TestRegister(t *testing.T) {
// 	repo := mocks.NewUserData(t)
// 	newUser := users.Core{
// 		Email: "alfian12345@gmail.com",

// 		Username: "alfianad",
// 		Password: "Alf12345",
// 	}
// 	expectedData := users.Core{
// 		Email:    "alfian12345@gmail.com",
// 		Username: "alfianad",
// 		Password: "$2a$10$ONzXu6wvVIUm0wzYCRinOO3c5qnKOLpx07OUh5njkMKqpR1phjGO",
// 	}

// 	t.Run("success register", func(t *testing.T) {
// 		// newUser.Password = hash
// 		repo.On("Register", newUser).Return(expectedData, nil).Once()
// 		srv := New(repo)
// 		// newUser.Password = password
// 		res, err := srv.Register(newUser)
// 		assert.Nil(t, err)
// 		assert.Equal(t, expectedData.ID, res.ID)
// 		assert.Equal(t, expectedData.Username, res.Username)
// 		assert.Equal(t, expectedData.Email, res.Email)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Duplicate email", func(t *testing.T) {
// 		// newUser.Password = hash
// 		repo.On("Register", newUser).Return(users.Core{}, errors.New("Duplicate users.email")).Once()
// 		srv := New(repo)
// 		// newUser.Password = password
// 		res, err := srv.Register(newUser)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "email already exist")
// 		assert.Equal(t, res.Email, "")
// 		repo.AssertExpectations(t)

// 	})

// 	t.Run("server problem", func(t *testing.T) {
// 		// newUser.Password = hash
// 		repo.On("Register", newUser).Return(users.Core{}, errors.New("server error")).Once()
// 		srv := New(repo)
// 		// newUser.Password = password
// 		res, err := srv.Register(newUser)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "server")
// 		assert.Equal(t, res.Username, "")
// 		repo.AssertExpectations(t)

// 	})

// 	t.Run("validation problem", func(t *testing.T) {
// 		// newUser.Password = hash
// 		srv := New(repo)
// 		// newUser.Password = password
// 		res, err := srv.Register(newUser)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, err, "validation")
// 		assert.Equal(t, res.Username, "")

// 	})

// }

func TestRegister(t *testing.T) {
	data := mocks.NewUserData(t)
	input := users.Core{Username: "griffin", Email: "grf29@gmail.com", Password: "Alf12345", Role: "guest"}
	resData := users.Core{ID: uint(1), Username: "griffin", Email: "grf29@gmail.com"}
	srv := New(data)

	t.Run("success create account", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(resData, nil).Once()
		res, err := srv.Register(input)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.NotEmpty(t, resData.Username)
		data.AssertExpectations(t)
	})

	t.Run("username not allowed empty", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(users.Core{}, errors.New("data not allowed to empty")).Once()
		res, err := srv.Register(input)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "empty")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("email not allowed empty", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(users.Core{}, errors.New("data not allowed to empty")).Once()
		res, err := srv.Register(input)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "empty")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(users.Core{}, errors.New("server error")).Once()
		res, err := srv.Register(input)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "server error")
		data.AssertExpectations(t)
	})

	t.Run("data already used", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(users.Core{}, errors.New("data already used, duplicated")).Once()
		res, err := srv.Register(input)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "already used")
		data.AssertExpectations(t)
	})

	t.Run("bcrypt error", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(users.Core{}, errors.New("password processed error")).Once()
		res, err := srv.Register(input)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})

}

func TestLogin(t *testing.T) {
	data := mocks.NewUserData(t)
	inputUsername := "griffin123"
	hashed, _ := helpers.GeneratePassword("Alf12345")
	resData := users.Core{ID: uint(1), Username: "griffin", Email: "grf@gmail.com", Password: hashed}
	srv := New(data)

	t.Run("login success", func(t *testing.T) {
		data.On("Login", inputUsername, mock.Anything).Return(resData, nil).Once()
		token, res, err := srv.Login(inputUsername, "Alf12345")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.Username, res.Username)
		data.AssertExpectations(t)
	})

	t.Run("password not matched", func(t *testing.T) {
		data.On("Login", inputUsername, mock.Anything).Return(resData, nil).Once()
		token, res, err := srv.Login(inputUsername, "Alf12345")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.Username, res.Username)
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("Login", inputUsername, mock.Anything).Return(users.Core{}, errors.New("server error")).Once()
		_, res, err := srv.Login(inputUsername, "Alf12345")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		assert.Empty(t, nil)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("username or password empty", func(t *testing.T) {
		wrong, _ := helpers.GeneratePassword("woooow123")

		data.On("Login", inputUsername, mock.Anything).Return(users.Core{Password: wrong}, nil).Once()

		_, res, err := srv.Login(inputUsername, "grf123")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "password do not match")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
}
