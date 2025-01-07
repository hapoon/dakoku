package main

import (
	"fmt"
	"ko07ga/dakoku/command"
	"ko07ga/dakoku/model"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
		Action:   dakoku,
		Commands: command.NewCommand(),
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func dakoku(ctx *cli.Context) (err error) {
	fmt.Println("打刻します")
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Fprintln(os.Stderr, "初期化が完了してないかもです")
		return
	}
	kintai := model.Kintai{}
	db.First(&kintai)
	fmt.Printf("%+v", kintai)

	return
}
