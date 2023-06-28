//go:build !test

package main

import (
	_ "embed"
	"fmt"
	"github.com/bridgewwater/temp-golang-cli-fast/command"
	"github.com/bridgewwater/temp-golang-cli-fast/command/subcommand_new"
	"github.com/bridgewwater/temp-golang-cli-fast/utils/pkgJson"
	"github.com/bridgewwater/temp-golang-cli-fast/utils/urfave_cli"
	"github.com/urfave/cli/v2"
	os "os"
)

//go:embed package.json
var packageJson string

func main() {
	pkgJson.InitPkgJsonContent(packageJson)
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = pkgJson.GetPackageJsonVersionGoStyle(false)
	app.Name = pkgJson.GetPackageJsonName()
	if pkgJson.GetPackageJsonHomepage() != "" {
		app.Usage = fmt.Sprintf("see: %s", pkgJson.GetPackageJsonHomepage())
	}
	app.Description = pkgJson.GetPackageJsonDescription()

	author := &cli.Author{
		Name:  pkgJson.GetPackageJsonAuthor().Name,
		Email: pkgJson.GetPackageJsonAuthor().Email,
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
		fmt.Printf("please see help as: %s --help\n", app.Name)
		os.Exit(2)
	}
	if err := app.Run(args); nil != err {
		fmt.Printf("run err: %v\n", err)
	}
}
