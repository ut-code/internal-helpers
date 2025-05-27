package fsutils

import (
	"fmt"
	"strconv"
	"strings"
)

// parses size string to int64
// example: 10kb -> 10_000
// example: 5MB -> 5_000_000
func ParseSize(size string) (int64, error) {
	size = strings.ToLower(size)
	if size == "" {
		return 0, fmt.Errorf("cannot parse empty string")
	}
	var unit int64
	switch  {
	case strings.HasSuffix(size, "kb"):
		unit = 1024
	case strings.HasSuffix(size, "mb"):
		unit = 1024 * 1024
	case strings.HasSuffix(size, "gb"):
		unit = 1024 * 1024 * 1024
	case strings.HasSuffix(size, "b"):
		unit = 1
	default:
		return 0, fmt.Errorf("unknown unit: %v", takeAfterInt(size))
	}
	n, err := takeWhileInt(size)
	if err != nil {
		return 0, err
	}
	return n * unit, nil
}

func takeAfterInt(str string) string {
	var i int
	for i = 0; i < len(str); i++ {
		if !strings.ContainsAny(string(str[i]), "0123456789") {
			break
		}
	}
	return string(str[i:])
}

func takeWhileInt(sizeStr string) (int64, error) {
	for i := 0; i < len(sizeStr); i++ {
		if !strings.ContainsAny(string(sizeStr[i]), "0123456789") {
			return strconv.ParseInt(string(sizeStr[:i]), 10, 64)
		}
	}
	return strconv.ParseInt(sizeStr, 10, 64)
}