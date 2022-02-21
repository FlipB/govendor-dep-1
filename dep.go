package sayhello

import "fmt"

// depVersion matches the git tag of the module for this dependency
const depVersion = "v3.0.0"

func versionInfo() string {
	return fmt.Sprintf("(from dep %s)", depVersion)
}

type Greeting struct {
	Phrase string
	Name   string
}

func (g Greeting) Greet() {
	msg := "Hello"
	if g.Phrase != "" {
		msg = g.Phrase
	}
	if g.Name != "" {
		msg += " " + g.Name
	}
	fmt.Printf("%s %s\n", msg, versionInfo())
}
