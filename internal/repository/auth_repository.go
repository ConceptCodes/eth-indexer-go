package repository

import (
	"gorm.io/gorm"

	"github.com/conceptcodes/eth-indexer-go/internal/models"
)

type AuthRepository interface {
	Create(user *models.Auth) error
	Delete(id string) error
	FindByToken(token string) (*models.Auth, error)
}

type GormAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &GormAuthRepository{db: db}
}

func (r *GormAuthRepository) Create(auth *models.Auth) error {
	return r.db.Create(auth).Error
}

func (r *GormAuthRepository) Delete(id string) error {
	return r.db.Delete(&models.Auth{}, id).Error
}

func (r *GormAuthRepository) FindByToken(token string) (*models.Auth, error) {
	var auth models.Auth

	if err := r.db.Where("token = ?", token).First(&auth).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}
