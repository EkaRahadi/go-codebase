package usecase

import (
	"context"
	"fmt"

	"github.com/EkaRahadi/go-codebase/internal/database"
	"github.com/EkaRahadi/go-codebase/internal/entity"
	apperror "github.com/EkaRahadi/go-codebase/internal/error" //alias
	"github.com/EkaRahadi/go-codebase/internal/repository"
)

type ExampleUsecase interface {
	ExampleUCFunc(ctx context.Context) (*entity.Dummy, error)
	ExampleUCTXFunc(ctx context.Context) (*entity.Dummy, error)
}

type exampleUsecase struct {
	exampleRepository repository.ExampleRepository
	transactor        database.Transactor
}

func NewExampleUsecase(exampleRepository repository.ExampleRepository, transactor database.Transactor) ExampleUsecase {
	return &exampleUsecase{
		exampleRepository: exampleRepository,
		transactor:        transactor,
	}
}

func (u *exampleUsecase) ExampleUCFunc(ctx context.Context) (*entity.Dummy, error) {
	res, err := u.exampleRepository.ExampleRepoFunc(ctx)
	if err != nil {
		return nil, apperror.NewServerError(fmt.Errorf("userUsecase.ExampleUCFunc: %w", err))
	}

	return res, nil
}

func (u *exampleUsecase) ExampleUCTXFunc(ctx context.Context) (*entity.Dummy, error) {
	var res *entity.Dummy
	var err error = nil
	err = u.transactor.Transaction(ctx, func(txCtx context.Context) error {
		res, err = u.exampleRepository.ExampleRepoFunc(txCtx)
		if err != nil {
			return apperror.NewServerError(fmt.Errorf("userUsecase.ExampleUCTXFunc: %w", err)) // this will trigger rollback
		}

		return nil // this will trigger commit
	})

	return res, err
}
