//go:build !test

package main

import (
	"fmt"
	"github.com/bridgewwater/temp-golang-cli-fast"
	"github.com/bridgewwater/temp-golang-cli-fast/command"
	"github.com/bridgewwater/temp-golang-cli-fast/command/subcommand_new"
	"github.com/bridgewwater/temp-golang-cli-fast/utils/pkgJson"
	"github.com/bridgewwater/temp-golang-cli-fast/utils/urfave_cli"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
	os "os"
	"time"
)

func main() {
	pkgJson.InitPkgJsonContent(temp_golang_cli_fast.PackageJson)
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
	app.Copyright = fmt.Sprintf("© 2022-%d %s", year, jsonAuthor.Name)
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

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("%s %s --help\n", color.Yellow.Render("please see help as:"), app.Name)
		os.Exit(2)
	}
	if err := app.Run(args); nil != err {
		color.Redf("cli err at %v\n", err)
	}
}
