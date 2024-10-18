// Package programming contains the functions for generating UUIDs.
package programming

import (
	"strings"

	"github.com/google/uuid"
)

// NewUUID generates a new UUID with the option to remove the hyphens.
func NewUUID(withoutHyphen bool) string {
	uuidWithHyphen := uuid.New()
	if withoutHyphen {
		return strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	}

	return uuidWithHyphen.String()
}
