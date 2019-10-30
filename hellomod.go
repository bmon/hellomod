package hellomod

import "fmt"

const version = "1.0.1"

func Hello() string {
	return fmt.Sprintf("Hello, modules, version: %s", version)
}
