package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n int
var arr [100000]int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if arr[i] > arr[i-(j*j)]+1 || arr[i] == 0 {
				arr[i] = arr[i-(j*j)] + 1
			}
		}
	}
	fmt.Fprint(bufout, arr[n])
}
