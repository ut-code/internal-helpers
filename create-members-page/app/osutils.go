package app

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
)

func mkDir(path string) error {
	return os.Mkdir(path, 0755)
}
func cp(src, dst string) error {
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

func mkDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return mkDir(path)
	}
	return nil
}

func mogrify(picturePath string) error {
	inExt := path.Ext(picturePath)
	withoutExt := strings.TrimSuffix(picturePath, inExt)
	outPath := withoutExt + ".avif"

	cmd := exec.Command("mogrify", "-format", "avif", "-quality", "90", "-resize", "800x", picturePath)
	if err := cmd.Run(); err != nil {
		return err
	}
	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		return fmt.Errorf("failed to create avif: %w", err)
	}
	// in case the original file is already an avif
	if picturePath != outPath {
		os.Remove(picturePath)
	}
	return nil
}
