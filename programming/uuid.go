// Package programming contains the functions for generating UUIDs.
package programming

import (
	"strings"

	"github.com/google/uuid"
)

// Service is the service for generating UUIDs.
type Service struct{}

// NewService creates a new UUID Service.
func NewService() Service {
	return Service{}
}

// NewUUID generates a new UUID with the option to remove the hyphens.
func (Service) NewUUID(withoutHyphen bool) string {
	uuidWithHyphen := uuid.New()
	if withoutHyphen {
		return strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	}

	return uuidWithHyphen.String()
}
