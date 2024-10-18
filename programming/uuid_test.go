package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUUIDWithHyphen(t *testing.T) {
	s := NewService()
	uuidWithHyphen := s.NewUUID(false)
	assert.Len(t, uuidWithHyphen, 36)
	assert.Contains(t, uuidWithHyphen, "-")
}

func TestNewUUIDWithoutHyphen(t *testing.T) {
	s := NewService()
	uuidWithoutHyphen := s.NewUUID(true)
	assert.Len(t, uuidWithoutHyphen, 32)
	assert.NotContains(t, uuidWithoutHyphen, "-")
}
