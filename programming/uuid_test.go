package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUUIDWithHyphen(t *testing.T) {
	uuidWithHyphen := NewUUID(false)
	assert.Len(t, uuidWithHyphen, 36)
	assert.Contains(t, uuidWithHyphen, "-")
}

func TestNewUUIDWithoutHyphen(t *testing.T) {
	uuidWithoutHyphen := NewUUID(true)
	assert.Len(t, uuidWithoutHyphen, 32)
	assert.NotContains(t, uuidWithoutHyphen, "-")
}
