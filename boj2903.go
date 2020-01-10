package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

var n int

func main() {
	defer bufout.Flush()
	fmt.Fscanf(bufin, "%d", &n)
	count := 2
	for i := 0; i < n; i++ {
		count += count - 1
	}
	fmt.Fprint(bufout, count*count)
}
