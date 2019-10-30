package hellomod

import "fmt"

const version = "2.0.0"

func Hello() string {
	return fmt.Sprintf("Hello, modules, version: %s", version)
}
