package main

import (
	"os"
	"log"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "redl"
	app.Usage = "A simple powerful cli tool for downloading courses inside Mixin ecosystem."
	app.Version = "1.1.0"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:	 "token",
			Aliases: []string{"t"},
			Usage:	 "collected after OAuth on the website (with Bearer prefix)",
		},
		&cli.StringFlag{
			Name:	 "base",
			Aliases: []string{"b"},
			Value:	 "xuexi-courses-api.songy.info",
			Usage:	 "base URL for downloading",
		},
		&cli.StringFlag{
			Name:    "dir",
			Aliases: []string{"d"},
			Usage:   "the output data directory",
		},
	}
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		{
			Name:	 "single",
			Aliases: []string{"s"},
			Usage:   "Download a single course",
			Action:  Single,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "dir",
					Aliases: []string{"d"},
					Usage:   "the output data directory",
				},
				&cli.StringFlag{
					Name:    "id",
					Aliases: []string{"i"},
					Usage:   "the course id to download",
				},
				&cli.StringFlag{
					Name:	 "token",
					Aliases: []string{"t"},
					Usage:	 "collected after OAuth on the website",
				},
				&cli.StringFlag{
					Name:	 "base",
					Aliases: []string{"b"},
					Value:	 "xuexi-courses-api.songy.info",
					Usage:	 "base URL for downloading",
				},
			},
		},
		{
			Name:	 "range",
			Aliases: []string{"r"},
			Usage:	 "Download all courses",
			Action:	 Range,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "range",
					Aliases: []string{"r"},
					Value:	 "0-1000",
					Usage:   "course range to download",
				},
			},
		},
		{
			Name:	 "all",
			Aliases: []string{"a"},
			Usage:	 "Download all courses",
			Action:	 All,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "dir",
					Aliases: []string{"d"},
					Value:	 ".",
					Usage:   "the data directory",
				},
				&cli.StringFlag{
					Name:	 "token",
					Aliases: []string{"t"},
					Usage:	 "collected after OAuth on the website",
				},
				&cli.StringFlag{
					Name:	 "base",
					Aliases: []string{"b"},
					Value:	 "xuexi-courses-api.songy.info",
					Usage:	 "base URL for downloading",
				},
			},
		},
	}
    
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
