package repository

import "subscription/internal/repository/db"

type Repository struct {
	DatabaseRepository *db.DatabaseRepository
}
