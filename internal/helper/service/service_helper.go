package service

import (
	"os"

	"github.com/EkaRahadi/go-codebase/internal/logger"
)

func GetInstanceID() string {
	if id := os.Getenv("HOSTNAME"); id != "" {
		return id //k8s podname as id
	}
	host, err := os.Hostname()
	if err != nil {
		logger.Log.Fatalw("Failed to get hostname", "error", err)
	}
	return host
}
