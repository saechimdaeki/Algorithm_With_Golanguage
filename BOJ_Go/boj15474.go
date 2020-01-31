package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

var n, a, b, c, d int

func min(x, y int) int {
	if x > y {
		return y
	} else if x == y {
		return x
	} else {
		return x
	}
}

var bsum, dsum int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n, &a, &b, &c, &d)
	for i := 1; i <= 1000; i++ {

		if n <= (a * i) {
			bsum += b
			break
		}
		bsum += b

	}
	for i := 1; i <= 1000; i++ {

		if n <= (c * i) {
			dsum += d
			break
		}
		dsum += d

	}
	fmt.Fprint(bufout, min(bsum, dsum))
}
