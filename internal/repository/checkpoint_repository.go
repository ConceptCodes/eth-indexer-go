package repository

import (
	"gorm.io/gorm"

	"github.com/conceptcodes/eth-indexer-go/internal/models"
)

type CheckpointRepository interface {
	Create(user *models.Checkpoint) error
	FindByID(id string) (*models.Checkpoint, error)
	FindLastBlock() (*models.Checkpoint, error)
}

type GormCheckpointRepository struct {
	db *gorm.DB
}

func NewCheckPointRepository(db *gorm.DB) CheckpointRepository {
	return &GormCheckpointRepository{db: db}
}

func (r *GormCheckpointRepository) Create(checkpoint *models.Checkpoint) error {
	return r.db.Create(checkpoint).Error
}

func (r *GormCheckpointRepository) FindByID(id string) (*models.Checkpoint, error) {
	var checkpoint models.Checkpoint
	if err := r.db.First(&checkpoint, id).Error; err != nil {
		return nil, err
	}
	return &checkpoint, nil
}

func (r *GormCheckpointRepository) FindLastBlock() (*models.Checkpoint, error) {
	var checkpoint models.Checkpoint
	if err := r.db.Last(&checkpoint).Error; err != nil {
		return nil, err
	}
	return &checkpoint, nil
}
