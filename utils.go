package io_base_service

import (
	"time"
)

// GetTimestamp returns the current timestamp in ISO 8601 format
func GetTimestamp() string {
	return time.Now().Format(time.RFC3339)
}
