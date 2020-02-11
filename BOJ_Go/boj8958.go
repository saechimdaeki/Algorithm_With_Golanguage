package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(bufin, &s)
		sum := 0
		cnt := 0
		for j := 0; j < len(s); j++ {
			if s[j] == 'X' {
				cnt = 0
			} else {
				cnt++
			}
			sum += cnt
		}
		fmt.Fprintln(bufout, sum)
	}
}
