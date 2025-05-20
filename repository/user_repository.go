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
	UpdateUser(userId string, user models.UserUpdateRequest) (models.User, error)
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

func (ur *userRepository) UpdateUser(userId string, userUpdate models.UserUpdateRequest) (models.User, error) {
	// 空文字も含めて更新するためmapで渡す
	updateData := map[string]any{
		"avatar_url": userUpdate.AvatarURL,
		"header_url": userUpdate.HeaderURL,
		"bio":        userUpdate.Bio,
	}
	if err := ur.db.Model(&models.User{}).Where("id=?", userId).Updates(updateData).Error; err != nil {
		return models.User{}, err
	}
	// 更新したユーザーを取得
	var user models.User
	if err := ur.db.Where("id=?", userId).First(&user).Error; err != nil {
		return models.User{}, err
	}
	// 更新したユーザーを返す
	return user, nil
}
