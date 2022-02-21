package sayhello

import "fmt"

// depVersion matches the git tag of the module for this dependency
const depVersion = "v1.1.0"

func say(greeting, name string) {
	fmt.Printf("%s%s (from %s)", greeting, name, depVersion)
}

func Hello() {
	say("Hello", "")
}

func GreetName(name string) {
	say("Hello ", name)
}
