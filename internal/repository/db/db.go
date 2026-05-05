package db

import (
	"fmt"
	"subscription/internal/config"

	"github.com/jmoiron/sqlx"
)

type DatabaseRepository struct {
	DB *sqlx.DB
}

func NewDatabaseInstance(envConf *config.Config) (*DatabaseRepository, error) {

	connectionString := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		envConf.Db.Host,
		envConf.Db.Port,
		envConf.Db.User,
		envConf.Db.Password,
		envConf.Db.Database)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	return &DatabaseRepository{DB: db}, nil
}

func (r *DatabaseRepository) Close() error {
	if r.DB != nil {
		return r.DB.Close()
	}
	return nil
}
