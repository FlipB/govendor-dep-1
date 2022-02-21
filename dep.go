package sayhello

import "fmt"

// depVersion matches the git tag of the module for this dependency
const depVersion = "v1.0.0"

func say(s string) {
	fmt.Printf("%s (from %s)", s, depVersion)
}

func Hello() {
	say("Hello")
}
