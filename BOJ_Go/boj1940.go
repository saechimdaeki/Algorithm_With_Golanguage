package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

var n, m, cnt int
var arr [15000]int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	fmt.Fscan(bufin, &m)
	//arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(bufin, &arr[i])
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i]+arr[j] == m {
				cnt++
			}
		}
	}
	fmt.Fprint(bufout, cnt)

}
