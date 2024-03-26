//go:build !test

package main

import (
	"fmt"
	"github.com/bridgewwater/temp-golang-cli-fast"
	"github.com/bridgewwater/temp-golang-cli-fast/cmd/cli"
	"github.com/bridgewwater/temp-golang-cli-fast/internal/pkg_kit"
	"github.com/gookit/color"
	os "os"
)

const (
	exitCodeCmdArgs = 2
)

func main() {
	pkg_kit.InitPkgJsonContent(temp_golang_cli_fast.PackageJson)

	app := cli.NewCliApp()

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("%s %s --help\n", color.Yellow.Render("please see help as:"), app.Name)
		os.Exit(exitCodeCmdArgs)
	}
	if err := app.Run(args); nil != err {
		color.Redf("cli err at %v\n", err)
	}
}
