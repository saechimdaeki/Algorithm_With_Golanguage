package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n int64
var answer int64

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	answer = ((n + 1) * n * (n - 1) / 2)

	fmt.Fprint(bufout, answer)
}
