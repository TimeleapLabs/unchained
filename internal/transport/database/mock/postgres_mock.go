package mock

import (
	"context"
	"testing"

	"github.com/KenshiTech/unchained/internal/ent"
	"github.com/KenshiTech/unchained/internal/transport/database"

	"github.com/peterldowns/pgtestdb"
)

type mockConnection struct {
	t  *testing.T
	db *ent.Client
}

func (m *mockConnection) GetConnection() *ent.Client {
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
			Port:       "5433",
			Options:    "sslmode=disable",
		},
		pgtestdb.NoopMigrator{},
	)

	var err error
	m.db, err = ent.Open("postgres", "postgresql://postgres:password@127.0.0.1:5433/unchained?sslmode=disable")
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
