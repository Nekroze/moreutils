package helpers

import (
	"bufio"
	"io"
	"os"
)

func StdinToString() string {
	if _, e := os.Stdin.Stat(); e != nil {
		panic(e)
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	return string(output)
}
