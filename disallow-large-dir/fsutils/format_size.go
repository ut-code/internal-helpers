package fsutils

import (
	"fmt"
	"log"
)

func FormatSize(size int64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
	)

	var suffix string
	var divider float64 // 1, KB, MB, GB
	var precision int // 0 -> 1.00, 1 -> 10.0, 2 -> 100
	var sizeIndex int // 0 -> B, 1 -> KB, 2 -> MB, 3 -> GB
	size_ := size
	for size_ > 9 {
		size_ /= 10
		precision++
	}
	for precision >= 3 && sizeIndex < 3 {
		precision -= 3
		sizeIndex++
	}
	switch sizeIndex {
	case 0:
		suffix = "B"
		// no need to divide: there's no decimal bits
		return fmt.Sprintf("%d %s", size, suffix)
	case 1:
		suffix = "KB"
		divider = float64(KB)
	case 2:
		suffix = "MB"
		divider = float64(MB)
	case 3:
		suffix = "GB"
		divider = float64(GB)
	}
	switch precision {
	case 2:
		return fmt.Sprintf("%.0f %s", float64(size)/divider, suffix)
	case 1:
		return fmt.Sprintf("%.1f %s", float64(size)/divider, suffix)
	case 0:
		return fmt.Sprintf("%.2f %s", float64(size)/divider, suffix)
	default:
		log.Fatalln("unreachable precision", precision)
		return ""
	}
}
