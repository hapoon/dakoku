package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	name    = "dakoku"
	version = "0.1.0"
)

func main() {
	app := &cli.App{
		Version:              version,
		Description:          "打刻をCLIで行うアプリケーションです",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "verbose", Usage: "verbose", Aliases: []string{"V"}},
		},
		Action: func(ctx *cli.Context) (err error) {
			fmt.Println("打刻しました")
			return
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
