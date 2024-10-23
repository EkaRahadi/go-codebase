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

func RegisterTokenRoutes(r *gin.Engine, gormWrapper *database.GormWrapper, vldtr utils.Validator, jwtUtil utils.JWTUtil) {
	userRepo := repository.NewUserRepository(gormWrapper)
	userUsecase := usecase.NewUserUsecase(userRepo)
	tokenHandler := ginhandler.NewTokenHandler(jwtUtil, userUsecase)

	r.POST("/token", middleware.JsonBody[dto.UserDummyRequest](vldtr), tokenHandler.GenerateAccessToken)
	token := r.Group("/token", middleware.AuthorizeRefreshJWT(jwtUtil))
	{
		token.POST("/refresh", tokenHandler.GenerateNewAccessToken)
	}
}
