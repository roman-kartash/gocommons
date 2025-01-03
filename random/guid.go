package random

import "github.com/google/uuid"

// Guid returns new guid.
func Guid() string {
	return uuid.Must(uuid.NewV7()).String()
}
