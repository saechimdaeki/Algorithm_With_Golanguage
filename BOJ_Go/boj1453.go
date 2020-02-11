package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var visited [101]bool
var n, cnt int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(bufin, &a)
		if visited[a] == true {
			cnt++
		} else {
			visited[a] = true
		}
	}
	fmt.Fprint(bufout, cnt)
}
