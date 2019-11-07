package hellomod

import "fmt"

const version = "2.3.0"

func Hello() string {
	return fmt.Sprintf("Hello, modules, version: %s", version)
}

func HelloWorld() string {
	return "Hello, world!"
}
