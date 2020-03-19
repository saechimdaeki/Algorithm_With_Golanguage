package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n int

func solve(n int) {
	if n == 0 {
		return
	}
	if n%2 == 0 {
		solve(-(n / 2))
		fmt.Print("0")
	} else {
		if n > 0 {
			solve(-(n / 2))
		} else {
			solve((-n + 1) / 2)
		}
		fmt.Print("1")
	}
}

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	if n == 0 {
		fmt.Fprintln(bufout, "0")
	} else {
		solve(n)
	}
}
