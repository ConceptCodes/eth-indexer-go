package repository

import (
	"gorm.io/gorm"

	"github.com/conceptcodes/eth-indexer-go/internal/models"
)

type UserRepository interface {
	Create(user *models.User) error
	Delete(id string) error
	Save(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByApiKey(apiKey string) (*models.User, error)
}

type GormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *GormUserRepository) Save(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *GormUserRepository) Delete(id string) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *GormUserRepository) FindByApiKey(apiKey string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("api_key = ? AND enabled = true AND email_verified > now() ", apiKey).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
