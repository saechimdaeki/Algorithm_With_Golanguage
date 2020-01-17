package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n, t, cnt, sum int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n, &t)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(bufin, &arr[i])
	}

	for i := 0; i < n; i++ {
		if arr[i] <= t {
			cnt++
		}
		t -= arr[i]
	}

	fmt.Fprint(bufout, cnt)

}
