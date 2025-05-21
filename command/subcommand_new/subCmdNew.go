package subcommand_new

import (
	"strings"

	"github.com/bridgewwater/temp-golang-cli-fast/command"
	"github.com/bridgewwater/temp-golang-cli-fast/constant"
	"github.com/bridgewwater/temp-golang-cli-fast/internal/cli_kit/urfave_cli"
	"github.com/bridgewwater/temp-golang-cli-fast/internal/d_log"
	"github.com/urfave/cli/v2"
)

const commandName = "new"

var commandEntry *NewCommand

type NewCommand struct {
	cliName  string
	version  string
	buildId  string
	homePage string

	Args cli.Args

	isDebug      bool
	execFullPath string
	runRootPath  string
	// TODO: remove it if not use
	PlatformConfig *constant.PlatformConfig
}

func (n *NewCommand) Exec() error {
	d_log.Debugf(
		"-> Exec cli [ %s ] by subCommand [ %s ], version %s buildID %s",
		n.cliName,
		commandName,
		n.version,
		n.buildId,
	)

	if n.isDebug {
		d_log.Verbosef("cli full path: %s", n.execFullPath)
		d_log.Verbosef("     run path: %s", n.runRootPath)
		d_log.Verbosef("     args len: %v", n.Args.Len())

		if n.Args.Len() > 0 {
			d_log.Verbosef("     args content: %s", strings.Join(n.Args.Slice(), " | "))
		}
	}

	return nil
}

func flag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "lib",
			Usage: "Use a library template",
			Value: false,
		},
		&cli.StringFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "Set the resulting package name, defaults to the directory name",
		},
	}
}

func withEntry(c *cli.Context) (*NewCommand, error) {
	d_log.Debugf("-> withEntry subCommand [ %s ]", commandName)

	if c.Bool("lib") {
		d_log.Info("new lib mode")
	}

	globalEntry := command.CmdGlobalEntry()

	return &NewCommand{
		cliName:  globalEntry.Name,
		version:  globalEntry.Version,
		buildId:  globalEntry.BuildId,
		homePage: globalEntry.HomePage,

		Args: c.Args(),

		isDebug:      globalEntry.Verbose,
		execFullPath: globalEntry.RootCfg.ExecFullPath,
		runRootPath:  globalEntry.RootCfg.RunRootPath,

		// todo: if not use platform config, remove this
		PlatformConfig: constant.BindPlatformConfig(c),
	}, nil
}

func action(c *cli.Context) error {
	d_log.Debugf("-> Sub Command action [ %s ] start", commandName)

	entry, err := withEntry(c)
	if err != nil {
		return err
	}

	commandEntry = entry

	return commandEntry.Exec()
}

func Command() []*cli.Command {
	return []*cli.Command{
		{
			Name:   commandName,
			Usage:  "",
			Action: action,

			// Flags: flag(),
			// todo: if not use platform config, remove this use method flag()
			// Flags: urfave_cli.UrfaveCliAppendCliFlag(flag(), command.HideGlobalFlag()),
			Flags: urfave_cli.UrfaveCliAppendCliFlag(flag(), constant.PlatformFlags()),
		},
	}
}
