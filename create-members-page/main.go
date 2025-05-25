package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/ut-code/internal-helpers/create-members-page/app"
)

func main() {
	cmd := &cli.Command{
		Name: "create-members-page",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "sheet",
				Aliases:  []string{"i", "s"},
				Usage:    "Path to Google Sheet File",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "pictures-directory",
				Aliases:  []string{"p", "pictures"},
				Usage:    "Directory with pictures",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "outdir",
				Aliases:  []string{"o", "out"},
				Usage:    "Output directory path",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "since",
				Usage:    "Only include members since this submission timestamp",
				Required: false,
			},
		},
		Action: app.Main,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatalln(err)
	}
}
