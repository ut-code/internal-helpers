package fsutils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSize(t *testing.T) {
	cases := []struct {
		name     string
		size     string
		expected int64
		err      error
	}{
		{
			name:     "empty string",
			size:     "",
			expected: 0,
			err:      fmt.Errorf("cannot parse empty string"),
		},
		{
			name:     "1B",
			size:     "1B",
			expected: 1,
			err:      nil,
		},
		{
			name:     "1KB",
			size:     "1KB",
			expected: 1024,
			err:      nil,
		},
		{
			name:     "1MB",
			size:     "1MB",
			expected: 1024 * 1024,
			err:      nil,
		},
		{
			name:     "5GB",
			size:     "5GB",
			expected: 5 * 1024 * 1024 * 1024,
			err:      nil,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			n, err := ParseSize(c.size)
			if c.err != nil {
				assert.Error(t, err)
				assert.Equal(t, c.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, c.expected, n)
			}
		})
	}
}

func TestTakeWhileInt(t *testing.T) {
	assert := assert.New(t)
	var i int64
	var err error

	i, err = takeWhileInt("1B")
	assert.NoError(err)
	assert.Equal(int64(1), i)
	i, err = takeWhileInt("10KB")
	assert.NoError(err)
	assert.Equal(int64(10), i)
	i, err = takeWhileInt("15MB")
	assert.NoError(err)
	assert.Equal(int64(15), i)
	i, err = takeWhileInt("10000GB")
	assert.NoError(err)
	assert.Equal(int64(10000), i)
}