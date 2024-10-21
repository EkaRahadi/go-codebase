package repository

import (
	"context"

	"github.com/EkaRahadi/go-codebase/internal/database"
	"github.com/EkaRahadi/go-codebase/internal/entity"
)

type ExampleRepository interface {
	ExampleRepoFunc(ctx context.Context) (*entity.Dummy, error)
}

type exampleRepository struct {
	db *database.GormWrapper
}

func NewExampleRepository(db *database.GormWrapper) ExampleRepository {
	return &exampleRepository{
		db: db,
	}
}

func (r *exampleRepository) ExampleRepoFunc(ctx context.Context) (*entity.Dummy, error) {
	res := &entity.Dummy{}

	// do query with gorm

	return res, nil
}
