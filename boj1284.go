package main

import (
	"bufio"
	"fmt"
	"os"
)

var s string
var n, width int
var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bufout.Flush()
	for {
		width = 0
		fmt.Fscan(bufin, &s)
		if s == "0" {
			break
		}
		for i := 0; i < len(s); i++ {
			width++
			if s[i] == '1' {
				width += 2
			} else if s[i] == '0' {
				width += 4
			} else {
				width += 3
			}
		}
		fmt.Fprintln(bufout, width+1)
	}
}
