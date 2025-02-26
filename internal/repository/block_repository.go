package repository

import (
	"gorm.io/gorm"

	"github.com/conceptcodes/eth-indexer-go/internal/models"
)

type BlockRepository interface {
	Create(user *models.Block) error
	Delete(id string) error
	FindByBlockNumber(blockNumber string) (*models.Block, error)
	Count() (int64, error)
	GetAll() ([]models.Block, error)
}

type GormBlockRepository struct {
	db *gorm.DB
}

func NewBlockRepository(db *gorm.DB) BlockRepository {
	return &GormBlockRepository{db: db}
}

func (r *GormBlockRepository) Create(block *models.Block) error {
	return r.db.Create(block).Error
}

func (r *GormBlockRepository) Delete(id string) error {
	return r.db.Delete(&models.Block{}, id).Error
}

func (r *GormBlockRepository) FindByBlockNumber(blockNumber string) (*models.Block, error) {
	var block models.Block
	err := r.db.Preload("Transactions").First(&block, "number = ?", blockNumber).Error
	return &block, err
}

func (r *GormBlockRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.Block{}).Count(&count).Error
	return count, err
}

func (r *GormBlockRepository) GetAll() ([]models.Block, error) {
	var blocks []models.Block
	err := r.db.Find(&blocks).Error
	return blocks, err
}
