package command

import "github.com/urfave/cli/v2"

type Commands []*cli.Command

type command struct {
	c Commands
}

type Command interface {
	Add(cp *cli.Command)
}

func (c *command) Add(cp *cli.Command) {
	c.c = append(c.c, cp)
	return
}

func NewCommand() Commands {
	c := command{}
	// コマンドはここで追加していく
	c.Add(NewInit())
	return c.c
}
