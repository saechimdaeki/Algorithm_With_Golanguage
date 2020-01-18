package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var a, b int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &a, &b)
	c := b - a
	fmt.Fprint(bufout, c, b)
}
