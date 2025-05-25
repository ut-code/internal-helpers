package app

import (
	"context"
	"fmt"
	"log"
	"path"
	"strings"

	"time"

	"github.com/urfave/cli/v3"
)

func Main(ctx context.Context, c *cli.Command) error {
	var since time.Time
	if c.String("since") != "" {
		var err error
		since, err = time.Parse("2006/01/02 15:04:05", c.String("since"))
		if err != nil {
			return fmt.Errorf("failed to parse since: %w", err)
		}
	} else {
		// since is empty, set it to the start of time (= no filtering applied)
		since = time.Time{}
	}
	ac := AppContext{
		PicturesDirectory: c.String("pictures-directory"),
		OutDirectory:      c.String("outdir"),
		Since:             since,
	}
	if !strings.HasSuffix(ac.PicturesDirectory, "/") {
		ac.PicturesDirectory += "/"
	}
	if !strings.HasSuffix(ac.OutDirectory, "/") {
		ac.OutDirectory += "/"
	}

	if err := mkDirIfNotExists(ac.OutDirectory); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	members, err := ParseFile(ac, c.String("sheet"))
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}
	for _, member := range members {
		if member.Metadata.Timestamp.Before(ac.Since) {
			continue
		}
		var err error
		var ext string
		var picdst string
		safeName := NameToFileName(member.Metadata.NameEn)
		memberDir := ac.OutDirectory + safeName + "/"
		err = mkDir(memberDir)
		if err != nil {
			err = fmt.Errorf("failed to create member directory: %w", err)
			goto end
		}
		WriteMember(ac, memberDir+"index.md", member)
		ext = path.Ext(member.Metadata.PicturePath)
		picdst = memberDir + safeName + ext
		err = cp(member.Metadata.PicturePath, picdst)
		if err != nil {
			err = fmt.Errorf("failed to copy picture: %w", err)
			goto end
		}
		if err = mogrify(picdst); err != nil {
			err = fmt.Errorf("failed to mogrify picture: %w", err)
			goto end
		}
	end:
		if err != nil {
			log.Println("error processing member:", member)
			log.Println(err)
		}
	}
	return nil
}
