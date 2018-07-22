package main

import (
	"github.com/lpisces/cli/cmd/boot"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "cmd"
	app.Usage = "cmd tool demo"

	app.Commands = []cli.Command{
		{
			Name:    "boot",
			Aliases: []string{"b"},
			Usage:   "cmd demo",
			Action:  boot.Run,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "debug, d",
					Usage: "debug switch",
				},
				cli.StringFlag{
					Name:  "config, c",
					Usage: "load config file",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
