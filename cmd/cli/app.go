package cli

import (
	"fmt"
	"github.com/bridgewwater/temp-golang-cli-fast/command"
	"github.com/bridgewwater/temp-golang-cli-fast/command/subcommand_new"
	"github.com/bridgewwater/temp-golang-cli-fast/internal/pkgJson"
	"github.com/bridgewwater/temp-golang-cli-fast/internal/urfave_cli"
	"github.com/bridgewwater/temp-golang-cli-fast/internal/urfave_cli/cli_exit_urfave"
	"github.com/urfave/cli/v2"
	"runtime"
	"time"
)

const (
	copyrightStartYear = "2023"
	defaultExitCode    = 1
)

func NewCliApp() *cli.App {
	cli_exit_urfave.ChangeDefaultExitCode(defaultExitCode)
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = pkgJson.GetPackageJsonVersionGoStyle(false)
	app.Name = pkgJson.GetPackageJsonName()
	if pkgJson.GetPackageJsonHomepage() != "" {
		app.Usage = fmt.Sprintf("see: %s", pkgJson.GetPackageJsonHomepage())
	}
	app.Description = pkgJson.GetPackageJsonDescription()
	year := time.Now().Year()
	jsonAuthor := pkgJson.GetPackageJsonAuthor()
	app.Copyright = fmt.Sprintf("© %s-%d %s by: %s, run on %s %s",
		copyrightStartYear, year, jsonAuthor.Name, runtime.Version(), runtime.GOOS, runtime.GOARCH)
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
