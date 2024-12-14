package main

import (
	"context"
	"log"
	"os"

	"github.com/averak/protobq/internal"
	"github.com/urfave/cli/v2"
)

func main() {
	if err := newCliApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newCliApp() *cli.App {
	res := cli.NewApp()
	res.Name = "protobq"
	res.Usage = "A tool for idempotent schema management in BigQuery using Protocol Buffers."
	res.Commands = []*cli.Command{
		{
			Name:  "apply",
			Usage: "apply schema",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "input",
					Usage:    "path to the .proto file",
					Required: true,
					Aliases:  []string{"i"},
				},
				&cli.StringFlag{
					Name:     "project-id",
					Usage:    "google cloud project id",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				err := internal.Apply(context.Background(), c.String("project-id"))
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "drop",
			Usage: "drop schema",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "input",
					Usage:    "path to the .proto file",
					Required: true,
					Aliases:  []string{"i"},
				},
				&cli.StringFlag{
					Name:     "project-id",
					Usage:    "google cloud project id",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				// TODO: implements me
				return nil
			},
		},
	}
	return res
}
