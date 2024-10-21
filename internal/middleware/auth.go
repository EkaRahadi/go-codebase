package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/EkaRahadi/go-codebase/internal/constants"
	"github.com/EkaRahadi/go-codebase/internal/dto"
	apperror "github.com/EkaRahadi/go-codebase/internal/error"
	"github.com/EkaRahadi/go-codebase/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizeJWT(jUtil utils.JWTUtil) gin.HandlerFunc {
	return func(c *gin.Context) {
		if jUtil.ShouldSkipValidation() {
			c.Next()
			return
		}

		authHeader := c.GetHeader(constants.HeaderName)
		s := strings.Split(authHeader, fmt.Sprintf("%v ", constants.Schema))
		authError := apperror.NewUnauthorizedError(errors.New("Unauthorized"))
		if len(s) < 2 {
			c.Error(authError)
			c.Abort()
			return
		}

		decodedToken := s[1]
		token, err := jUtil.ValidateAccessToken(decodedToken)
		if err != nil || !token.Valid {
			c.Error(authError)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Error(authError)
			c.Abort()
			return
		}

		userJson, _ := json.Marshal(claims["user"])
		var user dto.AccessUserJWT
		err = json.Unmarshal(userJson, &user)
		if err != nil {
			c.Error(authError)
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()

	}
}

func AuthorizeRefreshJWT(jUtil utils.JWTUtil) gin.HandlerFunc {
	return func(c *gin.Context) {
		if jUtil.ShouldSkipValidation() {
			c.Next()
		}

		authHeader := c.GetHeader(constants.HeaderName)
		s := strings.Split(authHeader, fmt.Sprintf("%v ", constants.Schema))
		authError := apperror.NewUnauthorizedError(errors.New("Unauthorized"))
		if len(s) < 2 {
			c.Error(authError)
			c.Abort()
			return
		}
		decodedToken := s[1]
		token, err := jUtil.ValidateRefreshToken(decodedToken)
		if err != nil || !token.Valid {
			c.Error(authError)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Error(authError)
			c.Abort()
			return
		}

		userJson, _ := json.Marshal(claims["user"])
		var user dto.RefreshUserJWT
		err = json.Unmarshal(userJson, &user)
		if err != nil {
			c.Error(authError)
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()

	}
}
