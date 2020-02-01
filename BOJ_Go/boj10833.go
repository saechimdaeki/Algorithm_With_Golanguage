package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var sum int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(bufin, &a, &b)
		sum += (b % a)

	}
	fmt.Fprint(bufout, sum)
}
