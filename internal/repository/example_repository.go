package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/EkaRahadi/go-codebase/internal/database"
	"github.com/EkaRahadi/go-codebase/internal/entity"
	"gorm.io/gorm"
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
	if err := r.db.Start(ctx).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil //gorm treat not found as error but sometimes no need to send error if no data found
		}

		return nil, fmt.Errorf("exampleRepository.ExampleRepoFuncWithTx: %w", err)
	}

	return res, nil
}
