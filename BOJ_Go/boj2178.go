package main

import (
	"fmt"
)

type pos struct {
	y, x, distance int
}

func main() {
	var N, M int
	var maze [105][105]int
	var visit [105][105]int
	fmt.Scanf("%d %d", &N, &M)
	// fmt.Println(N, M)

	for i := 1; i <= N; i++ {
		var t string
		fmt.Scanf("%s", &t)
		// fmt.Println(len(t))
		for j := 1; j <= len(t); j++ {
			maze[i][j] = int(t[j-1] - '0')
		}
	}

	q := make([]pos, 0)
	q = append(q, pos{1, 1, 1})
	for len(q) != 0 {
		p := q[0]
		// fmt.Printf("%d %d\n", p.y, p.x)
		if p.y == N && p.x == M {
			fmt.Printf("%d", p.distance)
			return
		}
		q = q[1:]
		dy := []int{-1, 0, 1, 0}
		dx := []int{0, 1, 0, -1}

		for k := 0; k < 4; k++ {
			newY := p.y + dy[k]
			newX := p.x + dx[k]
			if 1 <= newY && newY <= N &&
				1 <= newX && newX <= M && maze[newY][newX] == 1 {
				if visit[newY][newX] == 0 {
					q = append(q, pos{newY, newX, p.distance + 1})
					visit[newY][newX] = 1
				}
			}
		}
	}
}
