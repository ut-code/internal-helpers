package main

import (
	"fmt"
	"github.com/ut-code/internal-helpers/lib/fsutils"
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
			fatalMessages = append(fatalMessages, fmt.Sprintf("[ERROR] directory %v is too large (%v > %v)", dir, fsutils.FormatSize(size), fsutils.FormatSize(flags.MaxSize)))
		} else {
			fmt.Printf("[OK] %v %v\n", fsutils.FormatSize(size), dir)
		}
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
