package utils_test

import (
	"testing"

	"github.com/EkaRahadi/go-codebase/internal/config"
	"github.com/EkaRahadi/go-codebase/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestHashAndSalt(t *testing.T) {
	authUtil := utils.NewAuthUtil(&config.Config{})
	password := "securePassword123"

	hashedPassword, err := authUtil.HashAndSalt(password)

	assert.NoError(t, err, "HashAndSalt should not return an error")
	assert.NotEmpty(t, hashedPassword, "Hashed password should not be empty")
	assert.NotEqual(t, password, hashedPassword, "Hashed password should not equal the original password")
}

func TestComparePassword(t *testing.T) {
	authUtil := utils.NewAuthUtil(&config.Config{})
	password := "securePassword123"
	hashedPassword, _ := authUtil.HashAndSalt(password) // Ensure no error during hash generation

	match := authUtil.ComparePassword(hashedPassword, password)

	assert.True(t, match, "ComparePassword should return true for matching password")
	nonMatchingPassword := "wrongPassword"
	match = authUtil.ComparePassword(hashedPassword, nonMatchingPassword)
	assert.False(t, match, "ComparePassword should return false for non-matching password")
}

func TestHashAndSalt_ErrorHandling(t *testing.T) {
	authUtil := utils.NewAuthUtil(&config.Config{})

	emptyPassword := ""
	hashedPassword, err := authUtil.HashAndSalt(emptyPassword)

	assert.NoError(t, err, "HashAndSalt should not return an error for an empty password")
	assert.NotEmpty(t, hashedPassword, "Hashed password for an empty input should not be empty")
}
