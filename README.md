[![ci](https://github.com/bridgewwater/temp-golang-cli-fast/workflows/ci/badge.svg?branch=main)](https://github.com/bridgewwater/temp-golang-cli-fast/actions/workflows/ci.yml)
[![TravisBuildStatus](https://api.travis-ci.com/bridgewwater/temp-golang-cli-fast.svg?branch=main)](https://travis-ci.com/bridgewwater/temp-golang-cli-fast)

[![go mod version](https://img.shields.io/github/go-mod/go-version/bridgewwater/temp-golang-cli-fast?label=go.mod)](https://github.com/bridgewwater/temp-golang-cli-fast)
[![GoDoc](https://godoc.org/github.com/bridgewwater/temp-golang-cli-fast?status.png)](https://godoc.org/github.com/bridgewwater/temp-golang-cli-fast)
[![goreportcard](https://goreportcard.com/badge/github.com/bridgewwater/temp-golang-cli-fast)](https://goreportcard.com/report/github.com/bridgewwater/temp-golang-cli-fast)

[![GitHub license](https://img.shields.io/github/license/bridgewwater/temp-golang-cli-fast)](https://github.com/bridgewwater/temp-golang-cli-fast)
[![codecov](https://codecov.io/gh/bridgewwater/temp-golang-cli-fast/branch/main/graph/badge.svg)](https://codecov.io/gh/bridgewwater/temp-golang-cli-fast)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/bridgewwater/temp-golang-cli-fast)](https://github.com/bridgewwater/temp-golang-cli-fast/tags)
[![GitHub release)](https://img.shields.io/github/v/release/bridgewwater/temp-golang-cli-fast)](https://github.com/bridgewwater/temp-golang-cli-fast/releases)

### cli tools to init project fast

```bash
$ curl -L --fail https://raw.githubusercontent.com/bridgewwater/temp-golang-cli-fast/main/temp-golang-cli-fast
# let temp-golang-cli-fast file folder under $PATH
$ chmod +x temp-golang-cli-fast
# see how to use
$ temp-golang-cli-fast -h
```

## for what

- this project used to cli with golang

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/bridgewwater/temp-golang-cli-fast)](https://github.com/bridgewwater/temp-golang-cli-fast/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## Features

- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## env

- minimum go version: go 1.19
- change `go 1.19`, `^1.19`, `1.19.13` to new go version

### libs

| lib                                 | version |
|:------------------------------------|:--------|
| https://github.com/stretchr/testify | v1.8.4  |
| https://github.com/sebdah/goldie    | v2.5.3  |
| https://github.com/gookit/color     | v1.5.3  |
| https://github.com/bar-counter/slog | v1.4.0  |
| https://github.com/urfave/cli/      | v2.23.7 |

- more libs see [go.mod](https://github.com/bridgewwater/temp-golang-cli-fast/blob/main/go.mod)

## usage

```bash
# install at ${GOPATH}/bin
$ go install -v github.com/bridgewwater/temp-golang-cli-fast/cmd/temp-golang-cli-fast@latest
# install version v1.0.0
$ go install -v github.com/bridgewwater/temp-golang-cli-fast/cmd/temp-golang-cli-fast@v1.0.0
```

- use this template, replace list below and add usage
    - `github.com/bridgewwater/temp-golang-cli-fast` to your package name
    - `bridgewwater` to your owner name
    - `temp-golang-cli-fast` to your project name

# dev

## depends

in go mod project

```bash
# warning use private git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "https://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q https://github.com/bridgewwater/temp-golang-cli-fast.git

# test depends see full version
$ go list -mod readonly -v -m -versions github.com/bridgewwater/temp-golang-cli-fast
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/bridgewwater/temp-golang-cli-fast | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## local dev

```bash
# It needs to be executed after the first use or update of dependencies.
$ make init dep
```

- test code

```bash
$ make test testBenchmark
```

add main.go file and run

```bash
# run and shell help
$ make devHelp

# run at CLI_VERBOSE=true
$ make dev

# run at ordinary mode
$ make run
```

- ci to fast check

```bash
# check style at local
$ make style

# run ci at local
$ make ci
```

### docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# more info see
$ make helpDocker
```

### EngineeringStructure

```
.
├── Dockerfile                     # ci docker build
├── Dockerfile.s6                  # local docker build
├── Makefile                       # make entry
├── README.md
├── build                          # build output
├── cmd
│     └── temp-golang-cli-fast     # command line main package install and dev entrance
│         ├── main.go                   # command line entry
│         └── main_test.go              # integrated test entry
├── command                        # command line package
│         ├── TestMain.go             # common entry in unit test package
│         ├── flag.go                 # global flag
│         ├── global.go               # global command
│         ├── global_test.go          # global command unit test
│         ├── golder_data_test.go     # unit test test data case
│         ├── init_test.go            # unit test initialization tool
│         └── subcommand_new          # subcommandPackage new
├── constant                       # constant package 
│         └── env.go                  # constant environment variable
├── doc                            # command line tools documentation
│         └── cmd.md
├── go.mod
├── go.sum
├── package.json                   # command line profile information
├── resource.go                    # embed resource 
├── utils                          # toolkit package
│         ├── env_kit                 # environment variables toolkit
│         ├── log                     # log toolkit
│         ├── pkgJson                 # package.json toolkit
│         └── urfave_cli              # urfave/cli toolkit
├── vendor
└── z-MakefileUtils                # make toolkit

```

### log

- cli log use [github.com/bar-counter/slog](https://github.com/bar-counter/slog)
    - open debug log by env `CLI_VERBOSE=true` or global flag `--verbose`

```go
package foo

func action(c *cli.Context) error {
	slog.Debug("SubCommand [ new ] start") // this not show at CLI_VERBOSE=false

	if c.Bool("lib") {
		slog.Info("new lib mode")
	}
	return nil
}
```
