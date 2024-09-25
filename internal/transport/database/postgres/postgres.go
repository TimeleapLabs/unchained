package postgres

import (
	"context"

	"github.com/TimeleapLabs/unchained/internal/model"

	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type connection struct {
	db *gorm.DB
}

func (c *connection) HealthCheck(_ context.Context) bool {
	if c.db != nil {
		return false
	}
	conn, err := c.db.DB()
	if err != nil {
		utils.Logger.With("Error", err).Error("Failed to get DB connection")
		return false
	}

	err = conn.Ping()
	if err != nil {
		utils.Logger.With("Error", err).Error("Failed to ping DB")
		return false
	}

	return true
}

func (c *connection) GetConnection() *gorm.DB {
	if c.db != nil {
		return c.db
	}

	if config.App.Postgres.URL == "" {
		return nil
	}

	var err error

	utils.Logger.Info("Connecting to PostgresSQL")

	c.db, err = gorm.Open(
		postgres.Open(config.App.Postgres.URL),
		&gorm.Config{
			Logger:         logger.Default.LogMode(logger.Warn),
			TranslateError: true,
		},
	)
	if err != nil {
		panic(err)
	}

	return c.db
}

func (c *connection) Migrate() {
	if c.db == nil {
		c.GetConnection()
	}

	err := c.db.AutoMigrate(
		&model.Proof{},
		&model.CorrectnessDataFrame{},
		&model.EventLogDataFrame{},
		&model.AssetPriceDataFrame{},
	)

	if err != nil {
		utils.Logger.With("Error", err).Error("Failed to migrate DB")
	}
}

func New() database.Database {
	return &connection{}
}
