package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/EkaRahadi/go-codebase/internal/database"
	"github.com/EkaRahadi/go-codebase/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindOneByUserId(ctx context.Context, userId uint64) (*entity.User, error)
}

type userRepository struct {
	db *database.GormWrapper
}

func NewUserRepository(db *database.GormWrapper) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindOneByUserId(ctx context.Context, userId uint64) (*entity.User, error) {
	res := &entity.User{}

	if err := r.db.Start(ctx).First(&res, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("userRepository.FindOneByUserId: %w", err)
	}

	return res, nil
}
