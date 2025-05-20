package usecase

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/illionillion/go-sns-api-learning/models"
	"github.com/illionillion/go-sns-api-learning/repository"
)

type IUserUsecase interface {
	Signup(user models.User) (models.UserResponse, error)
	Login(user models.User) (string, error)
	GetUser(userId string) (models.UserResponse, error)
	UpdateUser(userId string, userUpdate models.UserUpdateRequest) (models.UserResponse, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) GetUser(userId string) (models.UserResponse, error) {
	user := models.User{}
	if err := uu.ur.GetUserById(&user, userId); err != nil {
		return models.UserResponse{}, err
	}
	return models.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
		HeaderURL: user.HeaderURL,
		Bio:       user.Bio,
	}, nil
}

func (uu *userUsecase) UpdateUser(userId string, userUpdate models.UserUpdateRequest) (models.UserResponse, error) {
	user, err := uu.ur.UpdateUser(userId, userUpdate)
	if err != nil {
		return models.UserResponse{}, err
	}
	return models.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
		HeaderURL: user.HeaderURL,
		Bio:       user.Bio,
	}, nil
}

// サインアップの処理
func (uu *userUsecase) Signup(user models.User) (models.UserResponse, error) {
	if err := user.SetPassword(user.Password); err != nil {
		return models.UserResponse{}, err
	}
	newUser := models.User{
		Email:    user.Email,
		Password: user.Password,
	}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return models.UserResponse{}, err
	}
	return models.UserResponse{
		ID:        newUser.ID,
		Email:     newUser.Email,
		AvatarURL: newUser.AvatarURL,
		HeaderURL: newUser.HeaderURL,
		Bio:       newUser.Bio,
	}, nil
}

// サインインの処理
func (uu *userUsecase) Login(user models.User) (string, error) {
	storedUser := models.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	if !storedUser.CheckPassword(user.Password) {
		return "", fmt.Errorf("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
