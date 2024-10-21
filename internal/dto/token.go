package dto

import "github.com/golang-jwt/jwt"

type AccessUserJWT struct {
	UserId   uint64 `json:"user_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type AccessTokenResponse struct {
	Token string        `json:"token"`
	User  AccessUserJWT `json:"user"`
}

type RefreshUserJWT struct {
	UserId uint64 `json:"user_id"`
}

type RefreshTokenResponse struct {
	Token string         `json:"token"`
	User  RefreshUserJWT `json:"user"`
}

type AccessTokenClaims struct {
	User AccessUserJWT `json:"user"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	User RefreshUserJWT `json:"user"`
	jwt.StandardClaims
}
