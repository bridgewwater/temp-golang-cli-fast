package command_test

import (
	"fmt"
	"github.com/sinlov-go/unittest-kit/env_kit"
)

// To use this lib package, just
// the SetOutput function when your application starts.
func Example() {
	// Basic usage of this lib

	fmt.Println("package name check")
	// Output:
	// package name check
}

// other different examples such as bar naming can be written like this
// When parsing, the first letter of the Example_ content will be automatically capitalized.
func Example_env_kit() {
	// mock EnvKeys
	// mock EnvKeys
	const keyEnvs = "ENV_KEYS"

	env_kit.SetEnvBool(keyEnvDebug, true)
	env_kit.SetEnvInt64(keyEnvCiNum, 2)
	env_kit.SetEnvStr(keyEnvCiKey, "foo")

	envDebug = env_kit.FetchOsEnvBool(keyEnvDebug, false)
	envCiNum = env_kit.FetchOsEnvInt(keyEnvCiNum, 0)
	envCiKey = env_kit.FetchOsEnvStr(keyEnvCiKey, "")

	fmt.Printf("envDebug %v\n", envDebug)
	fmt.Printf("envCiNum %v\n", envCiNum)
	fmt.Printf("envCiKey %v\n", envCiKey)
	// Output:
	// envDebug true
	// envCiNum 2
	// envCiKey foo
}
