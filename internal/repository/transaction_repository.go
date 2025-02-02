package repository

import (
	"gorm.io/gorm"

	"github.com/conceptcodes/eth-indexer-go/internal/models"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	FindByHash(hash string) (*models.Transaction, error)
	CreateAll(transactions []*models.Transaction) error
}

type GormTransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &GormTransactionRepository{db: db}
}

func (r *GormTransactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *GormTransactionRepository) FindByHash(id string) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Where("hash = ?", id).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *GormTransactionRepository) CreateAll(transactions []*models.Transaction) error {
	return r.db.CreateInBatches(transactions, len(transactions)).Error
}
