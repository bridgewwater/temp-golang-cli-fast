//go:build !test

package main

import (
	"fmt"
	os "os"

	temp_golang_cli_fast "github.com/bridgewwater/temp-golang-cli-fast"
	"github.com/bridgewwater/temp-golang-cli-fast/cmd/cli"
	"github.com/bridgewwater/temp-golang-cli-fast/constant"
	"github.com/bridgewwater/temp-golang-cli-fast/internal/cli_kit/pkg_kit"
	"github.com/bridgewwater/temp-golang-cli-fast/internal/d_log"
	"github.com/gookit/color"
)

const (
	// exitCodeCmdArgs SIGINT as 2.
	exitCodeCmdArgs = 2
)

//nolint:gochecknoglobals
var (
	// Populated by goreleaser during build.
	version    = "unknown"
	rawVersion = "unknown"
	buildID    string
	commit     = "?"
	date       = ""
)

func init() {
	if buildID == "" {
		buildID = "unknown"
	}
}

func main() {
	d_log.SetLogLineDeep(d_log.DefaultExtLogLineMaxDeep)
	pkg_kit.InitPkgJsonContent(temp_golang_cli_fast.PackageJson)

	bdInfo := pkg_kit.NewBuildInfo(
		pkg_kit.GetPackageJsonName(),
		pkg_kit.GetPackageJsonDescription(),
		version,
		rawVersion,
		buildID,
		commit,
		date,
		pkg_kit.GetPackageJsonAuthor().Name,
		constant.CopyrightStartYear,
	)

	app := cli.NewCliApp(bdInfo)

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("%s %s --help\n", color.Yellow.Render("please see help as:"), app.Name)
		os.Exit(exitCodeCmdArgs)
	}

	if err := app.Run(args); nil != err {
		color.Redf("cli err at %v\n", err)
		os.Exit(exitCodeCmdArgs)
	}
}
