package id

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	done := make(map[ID]bool)
	for i := 0; i < 1000; i++ {
		id := New()
		assert.False(t, done[id])
		assert.NotEqual(t, Nil, id)
		done[id] = true
	}
}

func TestParseID(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		parsed, err := ParseID("f47ac10b-58cc-4372-a567-0e02b2c3d479")
		assert.NoError(t, err)
		assert.NotEqual(t, Nil, parsed)
	})
	t.Run("wrong format", func(t *testing.T) {
		parsed, err := ParseID("f47ac10b-wrong-wrong-a567-0e02b2c3d479")
		assert.Error(t, err)
		assert.Equal(t, Nil, parsed)
	})
}

func TestID_String(t *testing.T) {
	parsed, err := ParseID("f47ac10b-58cc-4372-a567-0e02b2c3d479")
	require.NoError(t, err)
	assert.Equal(t, "f47ac10b-58cc-4372-a567-0e02b2c3d479", parsed.String())
}
