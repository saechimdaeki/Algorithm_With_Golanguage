package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n, m int64

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n, &m)
	sum := (n - m)
	if sum < 0 {
		fmt.Fprint(bufout, -sum)
	} else {
		fmt.Fprint(bufout, sum)
	}
}
