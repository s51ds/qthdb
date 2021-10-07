package main

import (
	"fmt"
	"github.com/s51ds/qthdb/db"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var app = cli.NewApp()

var cliFlags = []cli.Flag{&cli.IntFlag{
	Name:  "month",
	Usage: "data are focused to month, 1 is January",
	Value: 1,
}}

func commands() {
	app.Commands = []*cli.Command{
		{
			Name:        "makescp",
			Aliases:     []string{"scp"},
			Usage:       "make SCP file",
			UsageText:   "this command creates N1MM SCP file for VHF contest",
			Description: "description",
			ArgsUsage:   "args usage",
			Flags:       cliFlags,
			Category:    "",
			Action: func(c *cli.Context) error {
				fmt.Println("command:", c.Command.Name)
				if err := db.Open("./app/db.gob"); err != nil {
					fmt.Println(err.Error())
					dir, _ := os.Getwd()
					fmt.Println("Working directory:", dir)

				}
				fmt.Println(db.NumberOfRows())

				return nil
			},
		},
	}
}
func main() {
	app.Name = ""
	app.Usage = ""

	app = &cli.App{
		Flags: cliFlags,
	}

	commands()

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
