package postgres

import "github.com/uptrace/bun"

type Repository struct {
	db *bun.DB
}

func New(db *bun.DB) (*Repository, error) {
	return &Repository{
		db: db,
	}, nil
}
