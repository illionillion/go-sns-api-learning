package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

/*
ユーザーのデータ
*/
type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Email      string    `json:"email" gorm:"unique"`
	Password   string    `json:"password"`
	AvatarURL  string    `json:"avatar_url"`
	HeaderURL  string    `json:"header_url"`
	Bio        string    `json:"bio"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Check if given password is match to account's password
func (u *User) CheckPassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass)) == nil
}

// Hash password and set it to account object
func (u *User) SetPassword(pass string) error {
	passwordHash, err := generatePasswordHash(pass)
	if err != nil {
		return fmt.Errorf("generate error: %w", err)
	}
	u.Password = passwordHash
	return nil
}

func generatePasswordHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hashing password failed: %w", err)
	}
	return string(hash), nil
}

/*
ユーザーのサインアップリクエスト
*/
type UserSignupRequest struct {
    Email    string `json:"email" validate:"required,email" example:"test@example.com"`
    Password string `json:"password" validate:"required,min=6" example:"password123"`
}

/*
ユーザーのログインリクエスト
*/
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"test@example.com"`
	Password string `json:"password" validate:"required,min=6" example:"password123"`
}

/*
APIで返すユーザーのデータ
*/
type UserResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
	HeaderURL string `json:"header_url"`
	Bio       string `json:"bio"`
}
