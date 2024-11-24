package usecase_test

import (
	"context"
	"testing"

	"github.com/EkaRahadi/go-codebase/internal/entity"
	"github.com/EkaRahadi/go-codebase/internal/usecase"
	"github.com/EkaRahadi/go-codebase/mocks"
	mockRepo "github.com/EkaRahadi/go-codebase/mocks/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExampleUCFunc(t *testing.T) {
	t.Run("Should return *entity.Dummy and no error", func(t *testing.T) {
		mockExampleRepo := mockRepo.NewExampleRepository(t)
		mockTransactor := mocks.NewTransactor(t)
		dummy := &entity.Dummy{
			Message: "Test Dummy",
		}
		// dummyRes := dummy.GenerateCourierDTO()
		mockExampleRepo.On("ExampleRepoFunc", mock.Anything).Return(dummy, nil)
		exampleUsecase := usecase.NewExampleUsecase(mockExampleRepo, mockTransactor)

		ctx := context.TODO()
		res, err := exampleUsecase.ExampleUCFunc(ctx)

		assert.Equal(t, dummy, res)
		assert.Nil(t, err)
	})
}

func TestExampleUCTXFunc(t *testing.T) {
	t.Run("Should return *entity.Dummy and no error", func(t *testing.T) {
		mockExampleRepo := mockRepo.NewExampleRepository(t)
		mockTransactor := mocks.NewTransactor(t)
		dummy := &entity.Dummy{
			Message: "Test Dummy",
		}
		mockExampleRepo.On("ExampleRepoFunc", mock.Anything).Return(dummy, nil)
		mockTransactor.On("Transaction", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			txFunc := args.Get(1).(func(context.Context) error)
			_ = txFunc(context.Background()) // Call the transaction function
		})
		exampleUsecase := usecase.NewExampleUsecase(mockExampleRepo, mockTransactor)

		ctx := context.TODO()
		res, err := exampleUsecase.ExampleUCTXFunc(ctx)

		assert.Equal(t, dummy, res)
		assert.Nil(t, err)
	})
}
