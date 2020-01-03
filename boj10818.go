// 최소 최대문제입니다

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bufout.Flush()
	var a int
	fmt.Fscanln(bufin, &a)
	vector := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Fscan(bufin, &vector[i])
	}
	sort.Ints(vector)
	fmt.Fprintln(bufout, vector[0], vector[a-1])

}
