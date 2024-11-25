package service_test

import (
	"os"
	"testing"

	"github.com/EkaRahadi/go-codebase/internal/helper/service"
	"github.com/go-playground/assert/v2"
)

func TestGetInstanceID_WithEnvVariable(t *testing.T) {
	expectedID := "mock-pod-id"
	os.Setenv("HOSTNAME", expectedID)
	defer os.Unsetenv("HOSTNAME")

	instanceID := service.GetInstanceID()

	assert.Equal(t, expectedID, instanceID)
}
