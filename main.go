package main

import (
	"fmt"
	"github.com/easonlin404/gin-swagger/gen"
	"github.com/urfave/cli"
	"os"
)

var framework string

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "framework,f",
			Value:       "gin",
			Usage:       "web framework for the swagger",
			Destination: &framework,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "create doc.go",
			Action: func(c *cli.Context) error {
				gen.New().Build()

			},
		},
		{
			Name:    "update",
			Aliases: []string{"a"},
			Usage:   "update doc.go",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	app.Run(os.Args)
}