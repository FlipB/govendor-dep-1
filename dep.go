package sayhello

import "fmt"

// depVersion matches the git tag of the module for this dependency
const depVersion = "v2.0.1"

func say(greeting, name string) {
	fmt.Printf("%s %s (from dep %s)", greeting, name, depVersion)
}

func Hello(name string) {
	Greet("Hello", name)
}

func Greet(phrase, name string) {
	say(phrase, name)
}
