package utils_test

import (
	"testing"

	apperror "github.com/EkaRahadi/go-codebase/internal/error"

	"github.com/EkaRahadi/go-codebase/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestUser struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required,password"`
}

func TestCustomValidator_ValidInput(t *testing.T) {
	user := &TestUser{
		Username: "testuser123",
		Password: "Test@1234",
	}
	validator := utils.NewCustomValidator()

	err := validator.Validate(user)

	// Assert that there is no error
	assert.NoError(t, err)
}

func TestCustomValidator_InvalidUsername(t *testing.T) {
	// Create an invalid username (contains invalid characters)
	user := &TestUser{
		Username: "testuser@123",
		Password: "Test@1234",
	}
	validator := utils.NewCustomValidator()

	err := validator.Validate(user)

	require.Error(t, err)

	validationErr, ok := err.(*apperror.ValidationError)
	require.True(t, ok)
	assert.Len(t, validationErr.GetDetails(), 1)
	assert.Equal(t, validationErr.GetDetails()[0].Message, "Username must contain only string and digit")
}

func TestCustomValidator_InvalidPassword(t *testing.T) {
	user := &TestUser{
		Username: "testuser123",
		Password: "password",
	}
	validator := utils.NewCustomValidator()

	err := validator.Validate(user)

	require.Error(t, err)
	validationErr, ok := err.(*apperror.ValidationError)
	require.True(t, ok)
	assert.Len(t, validationErr.GetDetails(), 1)
	assert.Equal(t, validationErr.GetDetails()[0].Message, "Password must have at least 8 characters, 1 symbol, 1 capital letter, and 1 number")
}

func TestCustomValidator_MissingUsername(t *testing.T) {
	user := &TestUser{
		Username: "",
		Password: "Test@1234",
	}
	validator := utils.NewCustomValidator()

	err := validator.Validate(user)

	require.Error(t, err)
	validationErr, ok := err.(*apperror.ValidationError)
	require.True(t, ok)
	assert.Len(t, validationErr.GetDetails(), 1)
	assert.Equal(t, validationErr.GetDetails()[0].Message, "username is a required field")
}

func TestCustomValidator_MissingPassword(t *testing.T) {
	user := &TestUser{
		Username: "testuser123",
		Password: "",
	}
	validator := utils.NewCustomValidator()

	err := validator.Validate(user)

	require.Error(t, err)
	validationErr, ok := err.(*apperror.ValidationError)
	require.True(t, ok)
	assert.Len(t, validationErr.GetDetails(), 1)
	assert.Equal(t, validationErr.GetDetails()[0].Message, "password is a required field")
}
