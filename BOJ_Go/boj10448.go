package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

var t int

var arr [45]int

func ureka(n1 int) int {
	for i := 0; i < 44; i++ {
		for j := 0; j < 44; j++ {
			for k := 0; k < 44; k++ {
				if arr[i]+arr[j]+arr[k] == n1 {
					return 123
				}
			}
		}
	}
	return 0
}

var n int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &t)
	for i := 1; i < 45; i++ {
		arr[i-1] = (i*i + i) / 2
	}

	for i := 0; i < t; i++ {
		fmt.Fscan(bufin, &n)
		if ureka(n) == 123 {
			fmt.Fprintln(bufout, "1")
		} else {
			fmt.Fprintln(bufout, "0")
		}
	}
}
