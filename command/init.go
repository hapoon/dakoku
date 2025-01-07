package command

import (
	"fmt"
	"ko07ga/dakoku/model"
	"os"

	"github.com/urfave/cli/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(ctx *cli.Context) (err error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Fprintln(os.Stderr, "初期化に失敗しました")
		return
	}
	//db.Exec("CREATE TABLE kintais (id integer not null primary key, app_name, start_at, end_at)")
	db.AutoMigrate(&model.Kintai{})

	fmt.Fprintf(os.Stdout, "初期化に成功しました")
	return
}

func NewInit() *cli.Command {
	return &cli.Command{
		Name:      "init",
		Usage:     "設定を初期化します",
		UsageText: "各種設定の初期化を実行します",
		Action:    Init,
	}
}
