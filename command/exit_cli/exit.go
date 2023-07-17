package exit_cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

const exitCodeDefault = 127

var exitCode = exitCodeDefault

func ChangeDefaultExitCode(code int) {
	exitCode = code
}

func ExitCodeF(format string, a ...any) cli.ExitCoder {
	return cli.Exit(fmt.Sprintf(format, a...), exitCode)
}

func ExitF(code int, format string, a ...any) cli.ExitCoder {
	return cli.Exit(fmt.Sprintf(format, a...), code)
}

func ExitCodeErr(err error) cli.ExitCoder {
	return cli.Exit(err.Error(), exitCode)
}

func ExitErr(code int, err error) cli.ExitCoder {
	return cli.Exit(err.Error(), code)
}

func ExitCodeErrF(err error, msg string) cli.ExitCoder {
	return cli.Exit(fmt.Sprintf("%s err: %s", msg, err), exitCode)
}

func ExitErrF(code int, err error, msg string) cli.ExitCoder {
	return cli.Exit(fmt.Sprintf("%s err: %s", msg, err), code)
}
