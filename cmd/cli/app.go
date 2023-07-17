package cli

import (
	"fmt"
	"github.com/bridgewwater/temp-golang-cli-fast/command"
	"github.com/bridgewwater/temp-golang-cli-fast/command/subcommand_new"
	"github.com/bridgewwater/temp-golang-cli-fast/constant"
	"github.com/bridgewwater/temp-golang-cli-fast/utils/pkgJson"
	"github.com/bridgewwater/temp-golang-cli-fast/utils/urfave_cli"
	"github.com/urfave/cli/v2"
	"time"
)

func NewCliApp() *cli.App {
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
	app.Copyright = fmt.Sprintf("© %s-%d %s", constant.CopyrightStartYear, year, jsonAuthor.Name)
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
