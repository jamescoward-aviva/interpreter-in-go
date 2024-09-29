package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/flexer"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		for _, tok := range flexer.Flex(line) {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
