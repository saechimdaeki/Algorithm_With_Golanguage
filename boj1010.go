package main

import (
	"bufio"
	"fmt"
	"os"
)

var arr [31][31]int

func solution(x int, y int) int {
	if x == y || y == 0 {
		return 1
	}
	if arr[x-1][y-1] == 0 {
		arr[x-1][y-1] = solution(x-1, y-1)
	}
	if arr[x-1][y] == 0 {
		arr[x-1][y] = solution(x-1, y)
	}

	return arr[x-1][y-1] + arr[x-1][y]
}

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bufout.Flush()

	var a, b, t int
	fmt.Scanln(&t)
	for j := 0; j < t; j++ {
		fmt.Fscan(bufin, &a, &b)
		fmt.Fprintln(bufout, solution(b, a))
	}
}
