package profile

import (
	"database/sql"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/profile"
)

type DBProfile interface {
	Create() (*profile.Profile, error)
}

type PGProfileDB struct {
	db *sql.DB
}

func NewPGProfileDB(db *sql.DB) *PGProfileDB {
	return &PGProfileDB{db: db}
}

func (pg *PGProfileDB) Create() (*profile.Profile, error) {
	return nil, nil
}
