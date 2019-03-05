package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Nekroze/moreutils/pkg/helpers"
)

var usage = "[<haystack>] <needle>"
var description = "Compares the <haystack> (given on via argument or piped to stdin) prefix against the <needle>, returning non 0 exit status if <haystack> does not start with <needle>."

func help() {
	fmt.Fprintf(os.Stderr, "Usage: %s %s\n\n%s\n", os.Args[0], usage, description)
}

func main() {
	var haystack string
	var needle string
	arglen := len(os.Args)

	switch arglen {
	case 2:
		haystack = helpers.StdinToString()
		needle = os.Args[1]
	case 3:
		haystack = os.Args[1]
		needle = os.Args[2]
	default:
		help()
		// if no args given, exit 0 otherwise more than 3 must have been given so it will be non 0
		os.Exit(arglen - 1)
	}

	mustHavePrefix(haystack, needle)
}

func mustHavePrefix(s, prefix string) {
	if !strings.HasPrefix(s, prefix) {
		os.Exit(1)
	}
}
