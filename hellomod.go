package hellomod

import "fmt"

const version = "2.1.0"

func Hello() string {
	return fmt.Sprintf("Hello, modules, version: %s", version)
}
