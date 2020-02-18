package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var queue []int
var star [101]int   //중요도
var seoson [101]int //서순
var visited [101]bool
var t, n, m int

func solve() int {
	order := 1
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		index := 0
		check := 0
	z:
		for {
			if index >= n {
				break
			}
			if index == cur || visited[index] == true {
				index++
				goto z
			}
			if star[cur] < star[index] {
				queue = append(queue, cur)
				check = -1
				break
			}
			index++
		}
		if check == 0 { //우선순위 제일큰
			visited[cur] = true
			seoson[cur] = order
			order += 1
		}
	}
	return seoson[m]
}

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(bufin, &n, &m)
		for j := 0; j < 100; j++ {
			visited[j] = false
		}
		for j := 0; j < n; j++ {
			fmt.Fscan(bufin, &star[j])
			queue = append(queue, j)
		}
		fmt.Fprintln(bufout, solve())
	}

}
