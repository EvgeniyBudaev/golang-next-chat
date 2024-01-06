package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
