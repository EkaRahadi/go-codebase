package utils_test

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/EkaRahadi/go-codebase/internal/config"
	"github.com/EkaRahadi/go-codebase/internal/dto"
	"github.com/EkaRahadi/go-codebase/internal/entity"
	"github.com/EkaRahadi/go-codebase/internal/utils"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var appConfig = &config.Config{
	JWTConfig: config.JWTConfig{
		AccessSecretKey:      "secretKey",
		RefreshSecretKey:     "refreshSecretKey",
		Issuer:               "testIssuer",
		AccessTokenLifespan:  time.Hour,
		RefreshTokenLifespan: 3 * time.Hour,
	},
}

func TestJWTUtilGenerateAccessToken(t *testing.T) {
	jwtUtil := utils.NewJWTUtil(appConfig)
	user := &entity.User{
		UserId:   123,
		FullName: "John Doe",
		Email:    "johndoe@example.com",
	}

	tokenResp, err := jwtUtil.GenerateAccessToken(user)

	require.NoError(t, err)
	assert.NotEmpty(t, tokenResp.Token)

	// Decode the token to check claims
	token, err := jwt.Parse(tokenResp.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte("secretKey"), nil
	})
	require.NoError(t, err)

	// Assert that the token's claims match the user's data
	claims, ok := token.Claims.(jwt.MapClaims)
	require.True(t, ok)
	userJson, _ := json.Marshal(claims["user"])
	var userToken dto.AccessUserJWT
	_ = json.Unmarshal(userJson, &userToken)
	assert.Equal(t, userToken.UserId, user.UserId)
	assert.Equal(t, userToken.FullName, user.FullName)
	assert.Equal(t, userToken.Email, user.Email)
}

func TestJWTUtilValidateAccessToken(t *testing.T) {
	jwtUtil := utils.NewJWTUtil(appConfig)

	// Create a valid token using the same secret key
	claims := &dto.AccessTokenClaims{
		User: dto.AccessUserJWT{
			UserId:   123,
			FullName: "John Doe",
			Email:    "johndoe@example.com",
		},
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secretKey"))
	require.NoError(t, err)

	validatedToken, err := jwtUtil.ValidateAccessToken(tokenString)

	require.NoError(t, err)
	assert.NotNil(t, validatedToken)
}

func TestJWTUtilValidateAccessToken_InvalidToken(t *testing.T) {
	jwtUtil := utils.NewJWTUtil(appConfig)
	invalidToken := "invalidTokenString"

	_, err := jwtUtil.ValidateAccessToken(invalidToken)

	require.Error(t, err)
}

func TestJWTUtilGenerateRefreshToken(t *testing.T) {
	jwtUtil := utils.NewJWTUtil(appConfig)
	user := &entity.User{
		UserId: 123,
	}

	tokenResp, err := jwtUtil.GenerateRefreshToken(user)

	require.NoError(t, err)
	assert.NotEmpty(t, tokenResp.Token)

	// Decode the token to check claims
	token, err := jwt.Parse(tokenResp.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte("refreshSecretKey"), nil
	})

	require.NoError(t, err)

	// Assert that the token's claims match the user's data
	claims, ok := token.Claims.(jwt.MapClaims)
	require.True(t, ok)
	userJson, _ := json.Marshal(claims["user"])
	var userRefresh dto.RefreshUserJWT
	_ = json.Unmarshal(userJson, &userRefresh)
	assert.Equal(t, userRefresh.UserId, user.UserId)
}

func TestJWTUtilValidateRefreshToken(t *testing.T) {
	jwtUtil := utils.NewJWTUtil(appConfig)
	claims := &dto.RefreshTokenClaims{
		User: dto.RefreshUserJWT{
			UserId: 123,
		},
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("refreshSecretKey"))
	require.NoError(t, err)

	validatedToken, err := jwtUtil.ValidateRefreshToken(tokenString)

	require.NoError(t, err)
	assert.NotNil(t, validatedToken)
}

func TestJWTUtilValidateRefreshToken_InvalidToken(t *testing.T) {
	jwtUtil := utils.NewJWTUtil(appConfig)
	invalidToken := "invalidTokenString"

	_, err := jwtUtil.ValidateRefreshToken(invalidToken)

	require.Error(t, err)
}

func TestJWTUtilShouldSkipValidation(t *testing.T) {
	appConfig.App.Environment = "testing"
	jwtUtil := utils.NewJWTUtil(appConfig)

	assert.True(t, jwtUtil.ShouldSkipValidation())

	appConfig.App.Environment = "production"

	assert.False(t, jwtUtil.ShouldSkipValidation())
}
