package utils

import (
	"errors"
	"log"
	"time"

	"github.com/EkaRahadi/go-codebase/internal/config"
	"github.com/EkaRahadi/go-codebase/internal/dto"
	"github.com/EkaRahadi/go-codebase/internal/entity"
	apperror "github.com/EkaRahadi/go-codebase/internal/error"
	"github.com/golang-jwt/jwt"
)

type JWTUtil interface {
	ShouldSkipValidation() bool
	GenerateAccessToken(user *entity.User) (*dto.AccessTokenResponse, error)
	GenerateRefreshToken(user *entity.User) (*dto.RefreshTokenResponse, error)
	ValidateAccessToken(encodedToken string) (*jwt.Token, error)
	ValidateRefreshToken(encodedToken string) (*jwt.Token, error)
}

type jwtUtilImpl struct {
	config *config.Config
}

func NewJWTUtil(c *config.Config) JWTUtil {
	return &jwtUtilImpl{
		config: c,
	}
}

func (j jwtUtilImpl) GenerateAccessToken(user *entity.User) (*dto.AccessTokenResponse, error) {
	claims := &dto.AccessTokenClaims{
		User: dto.AccessUserJWT{
			UserId:   user.UserId,
			FullName: user.FullName,
			Email:    user.Email,
		},
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    j.config.JWTConfig.Issuer,
			ExpiresAt: time.Now().Add(j.config.JWTConfig.AccessTokenLifespan).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	hmacSecret := j.config.JWTConfig.AccessSecretKey
	tokenString, err := token.SignedString([]byte(hmacSecret))
	if err != nil {
		return nil, err
	}

	return &dto.AccessTokenResponse{
		Token: tokenString,
		User: dto.AccessUserJWT{
			UserId:   user.UserId,
			FullName: user.FullName,
			Email:    user.Email,
		},
	}, nil
}

func (j jwtUtilImpl) GenerateRefreshToken(user *entity.User) (*dto.RefreshTokenResponse, error) {
	claims := &dto.RefreshTokenClaims{
		User: dto.RefreshUserJWT{
			UserId: user.UserId,
		},
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    j.config.JWTConfig.Issuer,
			ExpiresAt: time.Now().Add(j.config.JWTConfig.RefreshTokenLifespan).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	hmacSecret := j.config.JWTConfig.RefreshSecretKey
	tokenString, err := token.SignedString([]byte(hmacSecret))
	if err != nil {
		return nil, err
	}

	return &dto.RefreshTokenResponse{
		Token: tokenString,
		User: dto.RefreshUserJWT{
			UserId: user.UserId,
		},
	}, nil
}

func (j jwtUtilImpl) ValidateAccessToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, apperror.NewUnauthorizedError(errors.New("Unauthorized"))
		}
		return []byte(j.config.JWTConfig.AccessSecretKey), nil
	})
}

func (j jwtUtilImpl) ValidateRefreshToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, apperror.NewUnauthorizedError(errors.New("Unauthorized"))
		}
		return []byte(j.config.JWTConfig.RefreshSecretKey), nil
	})
}

func (j jwtUtilImpl) ShouldSkipValidation() bool {
	if j.config.App.Environment == "testing" {
		log.Println("disable JWT authorization on testing env")
		return true
	}

	return false
}
