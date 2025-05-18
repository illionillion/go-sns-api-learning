package repository

import (
	"github.com/illionillion/go-sns-api-learning/models"
	"gorm.io/gorm"
)

// ユーザーリポジトリのインターフェース
type IUserRepository interface {
	GetUserByEmail(user *models.User, email string) error
	CreateUser(user *models.User) error
	GetUserById(user *models.User, userId string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserById(user *models.User, userId string) error {
	// ユーザーをIDで検索、ヒットすればnil、しなければerrorを返す
	if err := ur.db.Where("id=?", userId).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserByEmail(user *models.User, email string) error {
	// ユーザーをメールアドレスで検索、ヒットすればnil、しなければerrorを返す
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *models.User) error {
	// userのデータを作成、成功すればnil、失敗すればerrorを返す
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
