package magick

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func Mogrify(picturePath string, format string) error {
	var ext string
	switch format {
	case "avif", "webp", "jpg", "jpeg", "png":
		ext = "." + format
	case ".avif", ".webp", ".jpg", ".jpeg", ".png":
		ext = format
	default:
		return fmt.Errorf("unknown format: expected png, jpg, webp or avif, got %v", format)
	}

	inExt := path.Ext(picturePath)
	withoutExt := strings.TrimSuffix(picturePath, inExt)
	outPath := withoutExt + ext

	cmd := exec.Command("mogrify", "-format", "avif", "-quality", "90", "-resize", "800x", picturePath)
	if err := cmd.Run(); err != nil {
		return err
	}
	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		return fmt.Errorf("failed to create avif: %w", err)
	}
	// in case the original file is already an avif
	if picturePath != outPath {
		if err := os.Remove(picturePath); err != nil {
			return err
		}
	}
	return nil
}
