package mock

import (
	"context"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/TimeleapLabs/unchained/internal/transport/database"

	"github.com/peterldowns/pgtestdb"
)

type mockConnection struct {
	t  *testing.T
	db *gorm.DB
}

func (m *mockConnection) Migrate() {}

func (m *mockConnection) GetConnection() *gorm.DB {
	if m.db != nil {
		return m.db
	}

	_ = pgtestdb.New(
		m.t,
		pgtestdb.Config{
			DriverName: "pgx",
			User:       "postgres",
			Password:   "password",
			Host:       "localhost",
			Port:       "5432",
			Options:    "sslmode=disable",
		},
		pgtestdb.NoopMigrator{},
	)

	var err error
	m.db, err = gorm.Open(
		postgres.Open("postgresql://postgres:password@127.0.0.1:5432/unchained?sslmode=disable"),
		&gorm.Config{
			Logger:         logger.Default.LogMode(logger.Warn),
			TranslateError: true,
		},
	)
	if err != nil {
		panic(err)
	}

	return m.db
}

func (m *mockConnection) HealthCheck(_ context.Context) bool {
	return true
}

func New(t *testing.T) database.Database {
	return &mockConnection{
		t: t,
	}
}
