package subcommand_new

import (
	"github.com/bar-counter/slog"
	"github.com/bridgewwater/temp-golang-cli-fast/command"
	"github.com/bridgewwater/temp-golang-cli-fast/utils/urfave_cli"
	"github.com/urfave/cli/v2"
)

func Command() []*cli.Command {
	urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())
	return []*cli.Command{
		{
			Name:   "new",
			Usage:  "new command",
			Action: action,
			Flags:  flag(),
		},
	}
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "lib",
			Usage: "Use a library template",
			Value: false,
		},
		&cli.StringFlag{
			Name:  "name",
			Usage: "Set the resulting package name, defaults to the directory name",
		},
	}
}

func action(c *cli.Context) error {
	slog.Debug("SubCommand [ new ] start")
	entry, err := withEntry(c)
	if err != nil {
		return err
	}
	newCommandEntry = entry
	return newCommandEntry.Exec()
}

func withEntry(c *cli.Context) (*NewCommand, error) {
	if c.Bool("lib") {
		slog.Info("new lib mode")
	}
	globalEntry := command.CmdGlobalEntry()
	return &NewCommand{
		isDebug: globalEntry.Verbose,
	}, nil
}

var newCommandEntry *NewCommand

type NewCommand struct {
	isDebug bool
}

func (n *NewCommand) Exec() error {
	return nil
}
