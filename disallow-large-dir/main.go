package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	flags, err := ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("no directories specified")
	}
	var fatalMessages []string
	for _, dir := range args {
		size, err := dirSize(dir)
		if err != nil {
			fatalMessages = append(fatalMessages, err.Error())
		}
		if size > flags.MaxSize {
			fatalMessages = append(fatalMessages, fmt.Sprintf("directory %v is too large (%v > %v)", dir, size, flags.MaxSize))
		}
		fmt.Printf("%v %v\n", size, dir)
	}
	if len(fatalMessages) > 0 {
		log.Fatal("\n" + strings.Join(fatalMessages, "\n"))
	}
}

func dirSize(dir string) (int64, error) {
	stat, err := os.Stat(dir)
	if err != nil {
		return 0, fmt.Errorf("directory / file %v does not exist", dir)
	}
	if stat.IsDir() {
		files, err := os.ReadDir(dir)
		if err != nil {
			return 0, fmt.Errorf("directory %v does not exist", dir)
		}
		var size int64
		for _, file := range files {
			s, err := dirSize(dir + "/" + file.Name())
			if err != nil {
				return 0, err
			}
			size += s
		}
		return size, nil
	} else {
		return stat.Size(), nil
	}
}