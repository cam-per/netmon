package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/urfave/cli"
)

var app = &cli.App{
	Commands: []cli.Command{
		cli.Command{
			Name: "up",
			Action: func(ctx *cli.Context) error {
				fmt.Print(humanize.Bytes(speedTx(ctx.Args().Get(0))))
				return nil
			},
		},
		cli.Command{
			Name: "down",
			Action: func(ctx *cli.Context) error {
				fmt.Print(humanize.Bytes(speedRx(ctx.Args().Get(0))))
				return nil
			},
		},
	},
}
