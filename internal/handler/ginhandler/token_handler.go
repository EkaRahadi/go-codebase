package ginhandler

import (
	"fmt"

	"github.com/EkaRahadi/go-codebase/internal/dto"
	"github.com/EkaRahadi/go-codebase/internal/error"
	"github.com/EkaRahadi/go-codebase/internal/helper/response"
	"github.com/EkaRahadi/go-codebase/internal/usecase"
	"github.com/EkaRahadi/go-codebase/internal/utils"
	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	jUtil       utils.JWTUtil
	userUsecase usecase.UserUsecase
}

func NewTokenHandler(jUtil utils.JWTUtil, userUsecase usecase.UserUsecase) *TokenHandler {
	return &TokenHandler{
		jUtil:       jUtil,
		userUsecase: userUsecase,
	}
}

func (h *TokenHandler) Refresh(c *gin.Context) {
	userJWT, isExist := c.Get("user")
	if !isExist {
		c.Error(error.NewTokenError())
		return
	}

	var newAccessToken interface{}
	switch userJWT := userJWT.(type) {
	case dto.RefreshUserJWT:
		var userId = userJWT.UserId

		// TODO: get user by userId - verifying user existance
		user, err := h.userUsecase.GetOneById(c, userId)
		if err != nil {
			c.Error(err)
			return
		}

		newAccessToken, err = h.jUtil.GenerateAccessToken(user)
		if err != nil {
			c.Error(err)
			return
		}
	default:
		c.Error(fmt.Errorf("unexpected type for userJWT: %T", userJWT))
		return
	}

	response.ResponseOKData(c, newAccessToken)
}