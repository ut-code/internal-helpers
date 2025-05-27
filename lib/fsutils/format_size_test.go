package fsutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)
func TestFormatSize(t *testing.T) {
	assert := assert.New(t)
	// Bytes do not get decimal places
	assert.Equal("0 B", FormatSize(0))
	assert.Equal("10 B", FormatSize(10))
	// decimal + digits is always 3
	assert.Equal("150 KB", FormatSize(150*KB))
	assert.Equal("12.0 MB", FormatSize(12*MB))
	assert.Equal("1.00 GB", FormatSize(1 *GB))
	assert.Equal("10.0 GB", FormatSize(10*GB))
	assert.Equal("100 GB", FormatSize(100*GB))
}