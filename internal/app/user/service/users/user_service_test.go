package service

import (
	"context"
	"errors"
	"github.com/bobby-back-dev/golang-crud/helper/crypto"
	"github.com/bobby-back-dev/golang-crud/helper/reqres/reqresuser"
	"github.com/bobby-back-dev/golang-crud/internal/app/user/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Create(ctx context.Context, user *models.User) (*models.User, error) {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepo) Login(ctx context.Context, username string) (*models.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func TestUserService_UserCreateService_Success(t *testing.T) {
	mockRepo := new(MockUserRepo)
	hashHelper := &crypto.Hash{}
	respHelper := &reqresuser.UserWebRes{}
	userService := NewUserService(mockRepo, hashHelper, respHelper)
	ctx := context.Background()

	request := reqresuser.UserRequestRegisOrUpdate{
		Username:     "testuser",
		PasswordHash: "password123",
		DisplayName:  "Test User",
	}

	mockRepo.On("Create", ctx, mock.AnythingOfType("*models.User")).
		Return(&models.User{
			ID:          1,
			Username:    request.Username,
			DisplayName: request.DisplayName,
			CreatedAt:   time.Now(),
		}, nil).
		Once()

	response, err := userService.UserCreateService(ctx, request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, request.Username, response.Username)
	assert.Equal(t, int64(1), response.ID)

	mockRepo.AssertExpectations(t)
}

func TestUserService_UserCreateService_RepoError(t *testing.T) {
	mockRepo := new(MockUserRepo)
	hashHelper := &crypto.Hash{}
	respHelper := &reqresuser.UserWebRes{}
	userService := NewUserService(mockRepo, hashHelper, respHelper)
	ctx := context.Background()

	// PERBAIKAN: Menambahkan DisplayName agar lolos validasi.
	request := reqresuser.UserRequestRegisOrUpdate{
		Username:     "testuser",
		PasswordHash: "password123",
		DisplayName:  "Another Test User",
	}

	expectedError := errors.New("database connection lost")
	// Pastikan nama method di .On() adalah "CreateUser".
	mockRepo.On("Create", ctx, mock.AnythingOfType("*models.User")).
		Return(nil, expectedError).
		Once()

	response, err := userService.UserCreateService(ctx, request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), expectedError.Error())

	mockRepo.AssertExpectations(t)
}

func TestUserService_UserCreateService_InvalidData(t *testing.T) {
	mockRepo := new(MockUserRepo)
	hashHelper := &crypto.Hash{}
	respHelper := &reqresuser.UserWebRes{}
	userService := NewUserService(mockRepo, hashHelper, respHelper)
	ctx := context.Background()

	request := reqresuser.UserRequestRegisOrUpdate{
		Username:     "testuser",
		PasswordHash: "",
	}

	response, err := userService.UserCreateService(ctx, request)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "invalid data", err.Error())

	// Verifikasi bahwa method CreateUser di repo TIDAK PERNAH dipanggil.
	mockRepo.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
}

func TestUserService_LoginUser_Sucsses(t *testing.T) {
	mockRepo := new(MockUserRepo)
	hashHelper := &crypto.Hash{}
	respHelper := &reqresuser.UserWebRes{}
	userService := NewUserService(mockRepo, hashHelper, respHelper)
	ctx := context.Background()

	plainPassword := "password123"
	hashed, err := hashHelper.HashPassword(plainPassword)
	loginRequest := &reqresuser.UserRequestLogin{Username: "testuser", PasswordHash: plainPassword}

	mockRepo.On("Login", ctx, "testuser").
		Return(&models.User{
			ID:           1,
			Username:     "testuser",
			PasswordHash: hashed,
		}, nil).Once()

	response, err := userService.LoginUser(ctx, loginRequest)

	assert.NoError(t, err)

	assert.NotNil(t, response)
	assert.Equal(t, int64(1), response.ID)
	assert.Equal(t, "testuser", response.Username)

	mockRepo.AssertExpectations(t)
}

func TestUserService_LoginUser_Error(t *testing.T) {
	mockRepo := new(MockUserRepo)
	hashHelper := &crypto.Hash{}
	respHelper := &reqresuser.UserWebRes{}
	userService := NewUserService(mockRepo, hashHelper, respHelper)
	ctx := context.Background()

	wrongPwd := "password123"
	correctPwd := "coba123"
	hashed, err := hashHelper.HashPassword(correctPwd)
	loginRequest := &reqresuser.UserRequestLogin{Username: "testuser", PasswordHash: wrongPwd}

	mockRepo.On("Login", ctx, "testuser").
		Return(&models.User{
			ID:           1,
			Username:     "testuser",
			PasswordHash: hashed,
		}, nil).Once()

	response, err := userService.LoginUser(ctx, loginRequest)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "password is invalid", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestUserService_LoginUser_UserNotFound(t *testing.T) {
	mockRepo := new(MockUserRepo)
	hashHelper := &crypto.Hash{}
	respHelper := &reqresuser.UserWebRes{}
	userService := NewUserService(mockRepo, hashHelper, respHelper)
	ctx := context.Background()

	loginRequest := &reqresuser.UserRequestLogin{
		Username: "testuser1", PasswordHash: "coba123",
	}

	mockRepo.On("Login", ctx, "testuser1").
		Return(nil, errors.New("username or password is wrong")).Once()

	response, err := userService.LoginUser(ctx, loginRequest)
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, err.Error(), "username or password is wrong")
	mockRepo.AssertExpectations(t)
}
