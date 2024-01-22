package repository

import (
	"github.com/rogudator/tuvebes-backend/internal/entity"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	// Use the power of layered architecture and actually connect to database instead of stroing all in-memory
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) CreateTuvebe(t entity.Tuvebe) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	q := `
	INSERT INTO tuvebes(date, time_from, time_to, title)
VALUES 
    ('$1','$2','$3','$4');
	`
	_, err = tx.Exec(q, t.Date, t.TimeFrom, t.TimeTo, t.Title)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *Repository) GetTuvebes() []entity.Tuvebe {
	q := "select date,time_from,time_to,title from tuvebes"
	var tuvebes []entity.Tuvebe
	r.DB.Select(&tuvebes, q)
	return tuvebes
}
