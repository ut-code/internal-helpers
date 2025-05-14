package main

import (
	"flag"
	"github.com/ut-code/internal-helper/disallow-large-dir/fsutils"
)

type Flags struct {
	MaxSize int64
}

func ParseFlags() (Flags, error) {
	var maxSize string
	flag.StringVar(&maxSize, "max", "5MB", "maximum allowed size of a directory")
	flag.Parse()

	size, err := fsutils.ParseSize(maxSize)
	if err != nil {
		return Flags{}, err
	}
	return Flags{MaxSize: size}, nil
}
