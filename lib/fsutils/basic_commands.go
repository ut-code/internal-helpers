package fsutils

import (
	"io"
	"os"
)

func Mkdir(path string) error {
	return os.Mkdir(path, 0755)
}
func Cp(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return nil
}

func MkdirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return Mkdir(path)
	}
	return nil
}
