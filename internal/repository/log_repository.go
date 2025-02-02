package repository

import (
	"gorm.io/gorm"

	"github.com/conceptcodes/eth-indexer-go/internal/models"
)

type EventLogRepository interface {
	Create(user *models.Event) error
	CreateAll(logs []*models.Event) error
	FindByTransactionHash(txHash string) (*[]models.Event, error)
}

type GormEventLogRepository struct {
	db *gorm.DB
}

func NewEventLogRepository(db *gorm.DB) EventLogRepository {
	return &GormEventLogRepository{db: db}
}

func (r *GormEventLogRepository) Create(log *models.Event) error {
	return r.db.Create(log).Error
}

func (r *GormEventLogRepository) CreateAll(logs []*models.Event) error {
	return r.db.CreateInBatches(logs, len(logs)).Error
}

func (r *GormEventLogRepository) FindByTransactionHash(txHash string) (*[]models.Event, error) {
	var events []models.Event
	err := r.db.Where("transaction_hash = ?", txHash).Find(&events).Error
	return &events, err
}
