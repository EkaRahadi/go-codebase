package constants

import "time"

const (
	Issuer               = "Storee"
	AccessTokenLifespan  = time.Minute * 5
	RefreshTokenLifespan = time.Hour * 24
)

const (
	Schema     = "Bearer"
	HeaderName = "Authorization"
)
