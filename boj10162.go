package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var a, b, c, t int

func main() {
	defer bufout.Flush() 
	fmt.Fscanf(bufin, "%d", &t)
	if t%10 != 0 {
		fmt.Fprint(bufout, "-1")
	} else {
		a = t / 300
		t %= 300
		b = t / 60
		t %= 60
		c = t / 10
		t %= 10
		fmt.Fprint(bufout, a, b, c)
	}

}
