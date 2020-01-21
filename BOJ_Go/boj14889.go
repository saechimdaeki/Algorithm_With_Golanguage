package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var start, link, n int
var arr [21][21]int
var visited [21]bool

func dft(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

var infinite = 123456789

func dfs(n1, n2 int) {
	if n1 == (n / 2) {
		start = 0
		link = 0
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if visited[i] == true && visited[j] == true {
					start += arr[i][j]
				}
				if visited[i] == false && visited[j] == false {
					link += arr[i][j]
				}
			}
		}
		tmp := dft(start - link)
		if tmp < infinite {
			infinite = tmp
		}
		return
	}
	for i := n2; i < n; i++ {
		visited[i] = true
		dfs(n1+1, i+1)
		visited[i] = false
	}
}

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Fscan(bufin, &arr[i][j])
		}
	}
	dfs(0, 1)
	fmt.Fprint(bufout, infinite)
}
