package services

import (
	"alta-cookit-be/features/users"
	"alta-cookit-be/mocks"
	"alta-cookit-be/utils/helpers"
	"errors"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	data := mocks.NewUserData(t)
	input := users.Core{Username: "griffin", Email: "grf29@gmail.com", Password: "Alf12345"}
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

func TestProfile(t *testing.T) {
	data := mocks.NewUserData(t)
	resData := users.Core{ID: 1, Username: "griffin", Bio: "I love cooking", Email: "grf@gmail.com", Role: "User"}
	srv := New(data)

	t.Run("success show profile", func(t *testing.T) {
		data.On("Profile", uint(1)).Return(resData, nil).Once()

		res, err := srv.Profile(uint(1))
		assert.Nil(t, err)
		assert.Equal(t, resData.Username, res.Username)
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("Profile", uint(1)).Return(users.Core{}, errors.New("query error, problem with server")).Once()

		res, err := srv.Profile(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		assert.Equal(t, users.Core{}, res)
		data.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	data := mocks.NewUserData(t)
	filePath := filepath.Join("..", "..", "..", "test.jpg")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	inputData := users.Core{ID: 1, Email: "pian@gmail.com", Username: "alpian", Bio: "i love cooking"}
	resData := users.Core{ID: 1, Email: "pian@gmail.com", Username: "alpian", Bio: "i love cooking"}

	t.Run("success updating account", func(t *testing.T) {
		data.On("Update", uint(1), mock.Anything).Return(resData, nil).Once()
		srv := New(data)
		res, err := srv.Update(uint(1), *imageTrueCnv, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		data.AssertExpectations(t)
	})

	t.Run("fail updating account", func(t *testing.T) {
		data.On("Update", uint(1), mock.Anything).Return(users.Core{}, errors.New("user not found")).Once()
		srv := New(data)
		res, err := srv.Update(uint(1), *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not registered")
		assert.Equal(t, users.Core{}, res)
		data.AssertExpectations(t)
	})
	t.Run("email duplicated", func(t *testing.T) {
		data.On("Update", uint(1), mock.Anything).Return(users.Core{}, errors.New("email duplicated")).Once()
		srv := New(data)
		res, err := srv.Update(uint(1), *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "email duplicated")
		assert.Equal(t, users.Core{}, res)
		data.AssertExpectations(t)
	})
	t.Run("account not registered", func(t *testing.T) {
		data.On("Update", uint(1), mock.Anything).Return(users.Core{}, errors.New("access denied")).Once()
		srv := New(data)
		res, err := srv.Update(uint(1), *imageTrueCnv, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "access denied")
		assert.Equal(t, users.Core{}, res)
		data.AssertExpectations(t)
	})
}

func TestDeactive(t *testing.T) {
	data := mocks.NewUserData(t)
	t.Run("deleting account successful", func(t *testing.T) {
		data.On("Deactive", uint(1)).Return(nil).Once()
		srv := New(data)

		err := srv.Deactive(uint(1))
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})
	t.Run("internal server error, account fail to deactive", func(t *testing.T) {
		data.On("Deactive", uint(1)).Return(errors.New("no user has deactive")).Once()
		srv := New(data)
		err := srv.Deactive(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})
}

func TestUpdatePassword(t *testing.T) {
	data := mocks.NewUserData(t)
	// inputData := users.Core{Password: "Alfian123", NewPassword: "Alterra123", PasswordConfirmation: "Alterra123"}
	inputData := users.Core{Password: "Alfian123", NewPassword: "Alterra1234", PasswordConfirmation: "Alterra1234"}
	// comparePass, _ := helpers.CheckPassword(inputData.Password)
	// hashPass, _ := helpers.GeneratePassword("Alterra123")
	// checkPass, _ := helpers.CheckPassword(inputData.PasswordConfirmation)
	resData := users.Core{ID: uint(1), Username: "griffin", Bio: "I love cooking", Email: "grf@gmail.com", Role: "User", Password: "Alfian123", NewPassword: "Alterra123", PasswordConfirmation: "Alterra123"}

	t.Run("success update password", func(t *testing.T) {
		data.On("Profile", mock.Anything).Return(resData, nil).Once()
		data.On("UpdatePassword", mock.Anything, mock.Anything).Return(nil).Once()
		srv := New(data)
		err := srv.UpdatePassword(uint(1), inputData)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

}

func TestSearchUser(t *testing.T) {
	data := mocks.NewUserData(t)
	filePath := filepath.Join("..", "..", "..", "test.jpg")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	resData := []users.Core{{
		ID:             1,
		Username:       "abu",
		Role:           "User",
		Bio:            "i love cook",
		ProfilePicture: imageTrueCnv.Filename,
	}}
	q := "alfian"

	t.Run("success get all users", func(t *testing.T) {
		data.On("SearchUser", uint(1), q).Return(resData, nil).Once()
		srv := New(data)
		res, err := srv.SearchUser(uint(1), q)
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		data.AssertExpectations(t)
	})

	t.Run("users not found", func(t *testing.T) {
		data.On("SearchUser", uint(1), q).Return([]users.Core{}, errors.New("users not found")).Once()
		srv := New(data)
		res, err := srv.SearchUser(uint(1), q)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, 0, len(res))
		data.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		data.On("SearchUser", uint(1), q).Return([]users.Core{}, errors.New("server problem")).Once()
		srv := New(data)
		res, err := srv.SearchUser(uint(1), q)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, 0, len(res))
		data.AssertExpectations(t)
	})
}

func TestListUserRequest(t *testing.T) {
	data := mocks.NewUserData(t)
	filePath := filepath.Join("..", "..", "..", "test.jpg")
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}

	resData := []users.Core{{
		ID:             1,
		Username:       "abu",
		Role:           "VerifiedUser",
		Bio:            "i love cook",
		ProfilePicture: imageTrueCnv.Filename,
		Approvement:    "Requested",
	}}

	t.Run("success get all users", func(t *testing.T) {
		data.On("ListUserRequest", uint(1)).Return(resData, nil).Once()
		srv := New(data)
		res, err := srv.ListUserRequest(uint(1))
		assert.Nil(t, err)
		assert.Equal(t, len(resData), len(res))
		data.AssertExpectations(t)
	})

	t.Run("users not found", func(t *testing.T) {
		data.On("ListUserRequest", uint(1)).Return([]users.Core{}, errors.New("users not found")).Once()
		srv := New(data)
		res, err := srv.ListUserRequest(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, 0, len(res))
		data.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		data.On("ListUserRequest", uint(1)).Return([]users.Core{}, errors.New("server problem")).Once()
		srv := New(data)
		res, err := srv.ListUserRequest(uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, 0, len(res))
		data.AssertExpectations(t)
	})
}

func TestUpgradeUser(t *testing.T) {
	data := mocks.NewUserData(t)

	inputData := users.Core{Approvement: ""}

	resData := users.Core{ID: 1, Approvement: "requested"}

	t.Run("success send your request to admin", func(t *testing.T) {
		data.On("UpgradeUser", uint(1), inputData).Return(resData, nil).Once()
		srv := New(data)
		_, err := srv.UpgradeUser(uint(1), inputData)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("access denied", func(t *testing.T) {
		data.On("UpgradeUser", uint(1), inputData).Return(resData, errors.New("access denied")).Once()
		srv := New(data)
		res, err := srv.UpgradeUser(uint(1), inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "access denied")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("email duplicated", func(t *testing.T) {
		data.On("UpgradeUser", uint(1), inputData).Return(resData, errors.New("email duplicated")).Once()
		srv := New(data)
		res, err := srv.UpgradeUser(uint(1), inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "email duplicated")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
}
