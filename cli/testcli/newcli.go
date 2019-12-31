package main

import (
	"fmt"
	"os"

	"github.com/releax/releax/cli"
)

func main() {
	app := cli.New(&cli.App{
		Name:    "sample",
		Version: 0.1,
		Release: "alpha",
		Desc:    "Sample cli app",
		Usage:   "[args] -flags",
		Authors: []cli.AuthorData{
			cli.AuthorData{
				Name:  "Manjeet Saini",
				Email: "itsmanjeet1998@gmail.com",
			},
		},

		Sub: []cli.SubCommand{
			cli.SubCommand{
				Use:   "arg",
				Usage: "sample argument",
				Func: func(indata []string) error {
					fmt.Println("args", indata)
					return nil
				},
			},
		},
	})

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
