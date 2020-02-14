package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n, l int

func main() {
	defer bufout.Flush()
	cnt := 1
	fmt.Fscan(bufin, &n, &l)
	vector := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(bufin, &vector[i])
	}
	sort.Ints(vector)
	tmp := vector[0]
	for i := 1; i < n; i++ {
		if vector[i]-tmp+1 > l {
			cnt++
			tmp = vector[i]
		}
	}
	fmt.Fprint(bufout, cnt)
}
