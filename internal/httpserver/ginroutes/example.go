package ginroutes

import (
	"github.com/EkaRahadi/go-codebase/internal/database"
	"github.com/EkaRahadi/go-codebase/internal/dto"
	"github.com/EkaRahadi/go-codebase/internal/handler/ginhandler"
	"github.com/EkaRahadi/go-codebase/internal/middleware"
	"github.com/EkaRahadi/go-codebase/internal/repository"
	"github.com/EkaRahadi/go-codebase/internal/usecase"
	"github.com/EkaRahadi/go-codebase/internal/utils"
	"github.com/gin-gonic/gin"
)

func RegisterExampleRoutes(r *gin.Engine, gormWrapper *database.GormWrapper, transactor database.Transactor, vldtr utils.Validator) {
	exampleRepo := repository.NewExampleRepository(gormWrapper)
	exampleUsecase := usecase.NewExampleUsecase(exampleRepo, transactor)
	exampleHandler := ginhandler.NewExampleHandler(exampleUsecase)
	// Example middleware jwt auth,

	r.POST("/example", middleware.JsonBody[dto.DummyRequest](vldtr), exampleHandler.ExampleHandlerFunc)
	r.GET("/example-with-tx", middleware.Query[dto.DummyRequestQuery](vldtr), exampleHandler.ExampleHandlerWithTxFunc)
	r.GET("/example/:example_id", middleware.Uri[dto.DummyRequestUri](vldtr), exampleHandler.ExampleHandlerFuncUri)
}
