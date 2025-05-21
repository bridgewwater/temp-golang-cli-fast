package cli

import (
	"github.com/bridgewwater/temp-golang-cli-fast/command"
	"github.com/bridgewwater/temp-golang-cli-fast/command/subcommand_new"
	"github.com/bridgewwater/temp-golang-cli-fast/internal/cli_kit/pkg_kit"
	"github.com/bridgewwater/temp-golang-cli-fast/internal/cli_kit/urfave_cli"
	"github.com/bridgewwater/temp-golang-cli-fast/internal/cli_kit/urfave_cli/cli_exit_urfave"
	"github.com/urfave/cli/v2"
)

const (
	// defaultExitCode SIGINT as 2.
	defaultExitCode = 2
)

func NewCliApp(bdInfo pkg_kit.BuildInfo) *cli.App {
	cli_exit_urfave.ChangeDefaultExitCode(defaultExitCode)

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = bdInfo.PgkNameString()
	app.Version = bdInfo.VersionString()

	if pkg_kit.GetPackageJsonHomepage() != "" {
		app.Usage = "see: " + pkg_kit.GetPackageJsonHomepage()
	}

	app.Description = pkg_kit.GetPackageJsonDescription()
	jsonAuthor := pkg_kit.GetPackageJsonAuthor()
	app.Copyright = bdInfo.String()
	author := &cli.Author{
		Name:  jsonAuthor.Name,
		Email: jsonAuthor.Email,
	}
	app.Authors = []*cli.Author{
		author,
	}

	flags := urfave_cli.UrfaveCliAppendCliFlag(command.GlobalFlag(), command.HideGlobalFlag())

	app.Flags = flags
	app.Before = command.GlobalBeforeAction
	app.Action = command.GlobalAction
	app.After = command.GlobalAfterAction

	var appCommands []*cli.Command
	appCommands = urfave_cli.UrfaveCliAppendCliCommand(appCommands, subcommand_new.Command())

	app.Commands = appCommands

	return app
}
