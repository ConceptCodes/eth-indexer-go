package models

import (
	"time"

	"gorm.io/gorm"
)

type Block struct {
	gorm.Model

	Number       uint64 `gorm:"primaryKey;uniqueIndex"`
	Hash         string `gorm:"uniqueIndex"`
	ParentHash   string
	Nonce        string
	Miner        string
	Size         uint64
	Timestamp    uint64
	TxHash       string
	Transactions []Transaction `gorm:"foreignKey:BlockNumber;references:Number"`
}

func (b *Block) SimpleBlock() SimpleBlock {
	return SimpleBlock{
		Number:       b.Number,
		Hash:         b.Hash,
		Nonce:        b.Nonce,
		ParentHash:   b.ParentHash,
		Size:         b.Size,
		Miner:        b.Miner,
		Timestamp:    b.Timestamp,
		Transactions: SimplifyTransactions(b.Transactions),
	}
}

type Transaction struct {
	gorm.Model

	Hash        string `gorm:"primaryKey;uniqueIndex"`
	BlockNumber uint64 `gorm:"index"`

	From     string
	To       string
	Value    string
	GasPrice string
	GasLimit uint64
	GasUsed  uint64
	Nonce    uint64
}

func (t *Transaction) SimpleTransaction() SimpleTransaction {
	return SimpleTransaction{
		Hash:        t.Hash,
		BlockNumber: t.BlockNumber,
		From:        t.From,
		To:          t.To,
		Value:       t.Value,
		GasPrice:    t.GasPrice,
		GasLimit:    t.GasLimit,
		GasUsed:     t.GasUsed,
		Nonce:       t.Nonce,
	}
}

func SimplifyTransactions(data []Transaction) []SimpleTransaction {
	transactions := make([]SimpleTransaction, len(data))
	for i, t := range data {
		transactions[i] = t.SimpleTransaction()
	}
	return transactions
}

type Event struct {
	gorm.Model

	LogIndex        uint   `gorm:"primaryKey;autoIncrement:false"`
	TransactionHash string `gorm:"index"`
	BlockNumber     uint64 `gorm:"index"`
	Address         string
	Data            string   `gorm:"type:jsonb"`
	Topics          []string `gorm:"-"`
}

func (e *Event) SimpleEvent() SimpleEvent {
	return SimpleEvent{
		LogIndex:        e.LogIndex,
		TransactionHash: e.TransactionHash,
		BlockNumber:     e.BlockNumber,
		Address:         e.Address,
		Data:            e.Data,
	}
}

func SimplifyEvents(data []Event) []SimpleEvent {
	events := make([]SimpleEvent, len(data))
	for i, e := range data {
		events[i] = e.SimpleEvent()
	}
	return events
}

type Checkpoint struct {
	gorm.Model

	LastBlock uint64 `gorm:"uniqueIndex"`
}

type User struct {
	gorm.Model

	Name          string `json:"name" gorm:"index"`
	Email         string `gorm:"uniqueIndex"`
	Password      string `json:"password"`
	ApiKey        string `gorm:"uniqueIndex"`
	Enabled       bool   `json:"enabled"`
	EmailVerified time.Time
}

type Auth struct {
	gorm.Model

	Email     string `gorm:"uniqueIndex"`
	Token     string `gorm:"uniqueIndex"`
	ExpiresAt time.Time
	UserID    uint `gorm:"index"`
}
