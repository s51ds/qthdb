package main

import (
	"errors"
	"fmt"
	bl "github.com/s51ds/qthdb/app"
	"github.com/s51ds/qthdb/db"
	"github.com/s51ds/qthdb/timing"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
	"time"
)

// https://cyruslab.net/2020/11/06/golang-writing-a-command-line-program-with-urfave-cli-package/

var app = cli.NewApp()

var cliFlags = []cli.Flag{&cli.IntFlag{
	Name:    "month",
	Aliases: []string{"m", "M"},
	Usage:   "data are focused to the month, 1 is January",
	Value:   9,
}, &cli.StringFlag{
	Name:        "",
	Aliases:     nil,
	Usage:       "",
	EnvVars:     nil,
	FilePath:    "",
	Required:    false,
	Hidden:      false,
	TakesFile:   false,
	Value:       "",
	DefaultText: "",
	Destination: nil,
	HasBeenSet:  false,
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
			Action: func(c *cli.Context) error {
				fmt.Println("command:", c.Command.Name)
				if err := db.Open("./db.gob"); err != nil {
					fmt.Println(err.Error())
					dir, _ := os.Getwd()
					fmt.Println("Working directory:", dir)
					return err

				}
				m := c.Int("month")
				if m < 1 || m > 12 {
					return errors.New("month JAN=1...DEC=12")
				}
				return bl.MakeN1mmScpFile(fmt.Sprintf("scp-%s-%s.txt", timing.ShortMonthNames[m-1], strconv.Itoa(time.Now().Year())), time.Month(m))
			},
		},
	}
}
func main() {
	app.Name = "qthdb"
	app.Usage = "Example: qthdb --month 9 scp"

	app = &cli.App{
		Flags: cliFlags,
	}

	commands()

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
