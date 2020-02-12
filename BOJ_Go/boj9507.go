package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var t int
var dynamic [70]int64

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &t)
	dynamic[0] = 1
	dynamic[1] = 1
	dynamic[2] = 2
	dynamic[3] = 4
	for i := 4; i < 70; i++ {
		dynamic[i] = dynamic[i-1] + dynamic[i-2] + dynamic[i-3] + dynamic[i-4]
	}
	for i := 0; i < t; i++ {
		var n int64
		fmt.Fscan(bufin, &n)
		fmt.Fprintln(bufout, dynamic[n])
	}
}
