package main

import (
	"fmt"
	"sort"
)

type pair struct {
	f, s int
}

type queue struct {
	a []pair
}

func (q *queue) push(n1, n2 int) {
	q.a = append(q.a, pair{n1, n2})
}

func (q *queue) pop() {
	if len(q.a) != 0 {
		q.a = q.a[1:]
	}
}

func (q *queue) empty() bool {
	b := false
	if len(q.a) == 0 {
		b = true
	}
	return b
}

func printA(a [][]int, num int) {
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
}

func bfs(m [][]int, y int, x int) ([][]int, int) {
	q := queue{}
	cnt := 0
	q.push(y, x)
	m[y][x] = 0
	dy := []int{1, 0, -1, 0}
	dx := []int{0, -1, 0, 1}
	for !q.empty() {
		cnt++
		cy := q.a[0].f
		cx := q.a[0].s
		for i := 0; i < 4; i++ {
			ny := dy[i] + cy
			nx := dx[i] + cx
			if ny < 0 || ny > len(m[0])-1 || nx < 0 || nx > len(m[0])-1 {
				continue
			}
			if m[ny][nx] == 1 {
				m[ny][nx] = 0
				q.push(ny, nx)
			}
		}
		q.pop()
	}
	return m, cnt
}

func main() {
	var num1 int
	var num2 string
	var temp int
	cnt := []int{}
	fmt.Scan(&num1)
	m := make([][]int, num1)
	for i := 0; i < num1; i++ {
		m[i] = make([]int, num1)
		fmt.Scan(&num2)
		for j := 0; j < num1; j++ {
			m[i][j] = int(num2[j] - '0')
		}
	}
	for i := 0; i < num1; i++ {
		for j := 0; j < num1; j++ {
			if m[i][j] == 1 {
				m, temp = bfs(m, i, j)
				cnt = append(cnt, temp)
			}
		}
	}
	sort.Ints(cnt)
	fmt.Println(len(cnt))
	for i := 0; i < len(cnt); i++ {
		fmt.Println(cnt[i])
	}
}
