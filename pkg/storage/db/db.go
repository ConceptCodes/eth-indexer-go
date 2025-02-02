package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/conceptcodes/eth-indexer-go/config"
	"github.com/conceptcodes/eth-indexer-go/internal/models"
	"github.com/rs/zerolog"
)

var DB *gorm.DB

type DBConfig struct {
	cfg    *config.Config
	logger *zerolog.Logger
	client gorm.DB
}

func NewDB(cfg *config.Config, logger *zerolog.Logger) *DBConfig {
	return &DBConfig{
		cfg:    cfg,
		logger: logger,
		client: gorm.DB{},
	}
}

func (db *DBConfig) ConnectDB() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		db.cfg.DBHost, db.cfg.DBPort, db.cfg.DBUser, db.cfg.DBPassword, db.cfg.DBName,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		db.logger.Fatal().Err(err).Msgf("Failed to connect to database: %v", err)
	}

	db.logger.Info().Msg("Connected to PostgreSQL database successfully")

	// Run migrations
	err = DB.AutoMigrate(
		&models.Block{},
		&models.Transaction{},
		&models.Checkpoint{},
		&models.Event{},
	)
	if err != nil {
		db.logger.Fatal().Msgf("Failed to migrate database schema: %v", err)
	}

	db.client = *DB
	db.logger.Debug().Msg("Database schema migrated successfully")
}

func (db *DBConfig) GetDB() *gorm.DB {
	return DB
}

func (db *DBConfig) CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		db.logger.Fatal().Err(err).Msg("Failed to close database connection")
	}

	err = sqlDB.Close()
	if err != nil {
		db.logger.Fatal().Err(err).Msg("Failed to close database connection")
	}

	db.logger.Info().Msg("Database connection closed successfully")
}

func (db *DBConfig) HeathCheck() error {
	return DB.Exec("SELECT 1").Error
}
