package repository

import "github.com/rogudator/tuvebes-backend/internal/entity"

type Repository struct {
	// Use the power of layered architecture and actually connect to database instead of stroing all in-memory
	DB []entity.Tuvebe
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateTuvebe(t entity.Tuvebe) {
	r.DB = append(r.DB, t)
}

func (r *Repository) GetTuvebes() []entity.Tuvebe {
	return r.DB
}