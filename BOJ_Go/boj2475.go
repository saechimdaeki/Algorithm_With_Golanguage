package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n1, n2, n3, n4, n5 int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n1, &n2, &n3, &n4, &n5)
	fmt.Fprint(bufout, ((n1*n1)+(n2*n2)+(n3*n3)+(n4*n4)+(n5*n5))%10)
}
