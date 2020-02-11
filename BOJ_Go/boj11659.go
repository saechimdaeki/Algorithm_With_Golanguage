package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n, m int
var a, b int
var sum int
var arr [100001]int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n, &m)
	for i := 1; i <= n; i++ {
		var tmp int
		fmt.Fscan(bufin, &tmp)
		arr[i] = arr[i-1] + tmp
	}
	//fmt.Fprint(bufout, "-------------------------")
	for i := 0; i < m; i++ {
		fmt.Fscan(bufin, &a, &b)
		if a == 0 {
			fmt.Fprintln(bufout, arr[b])
		} else {
			fmt.Fprintln(bufout, arr[b]-arr[a-1])
		}
	}
}
