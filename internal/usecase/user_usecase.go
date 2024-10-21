package usecase

import (
	"context"
	"fmt"

	"github.com/EkaRahadi/go-codebase/internal/constants"
	"github.com/EkaRahadi/go-codebase/internal/entity"
	apperror "github.com/EkaRahadi/go-codebase/internal/error"
	"github.com/EkaRahadi/go-codebase/internal/repository"
)

type UserUsecase interface {
	GetOneById(ctx context.Context, userId uint64) (*entity.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) GetOneById(ctx context.Context, userId uint64) (*entity.User, error) {
	res, err := u.userRepository.FindOneByUserId(ctx, userId)
	if err != nil {
		return nil, apperror.NewServerError(fmt.Errorf("userUsecase.GetOneById: %w", err))
	}

	if res == nil {
		return nil, apperror.NewDataNotFoundError(constants.EntityUser, userId)
	}

	return res, nil
}
