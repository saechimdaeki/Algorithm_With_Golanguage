package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, m, num1, num2, x, y int
var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var grid [101][101]int
var visited [101]bool
var count [101]int

func bfs(start int) {
	queue := make([]int, 0)
	queue = append(queue, start)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for i := 1; i <= n; i++ {
			if grid[cur][i] == 1 && visited[i] == false {
				visited[i] = true
				count[i] = count[cur] + 1
				queue = append(queue, i)
			}
		}
	}
}
func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	fmt.Fscan(bufin, &num1, &num2)
	fmt.Fscan(bufin, &m)
	var a1, a2 int
	for i := 1; i <= m; i++ {
		fmt.Fscan(bufin, &a1, &a2)
		grid[a1][a2] = 1
		grid[a2][a1] = 1
	}
	bfs(num1)
	if count[num2] != 0 {
		fmt.Fprintln(bufout, count[num2])
	} else {
		fmt.Fprintln(bufout, "-1")
	}
}
