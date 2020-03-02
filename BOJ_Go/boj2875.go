package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

var n, m, k int

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n, &m, &k)
	answer := 0
	answer = min((n+m-k)/3, n/2)
	fmt.Fprint(bufout, min(answer, m))
}
