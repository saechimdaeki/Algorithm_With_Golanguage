package main

import "fmt"

var data [2501][2501]int
var check [2501][2501]int
var n int
var m int

func clear() {
	for i := 0; i < 2501; i++ {
		for j := 0; j < 2501; j++ {
			data[i][j] = 0
			check[i][j] = 0
		}
	}
}

func dfs(y int, x int) {
	dy := [4]int{0, 0, -1, 1}
	dx := [4]int{-1, 1, 0, 0}
	for i := 0; i < 4; i++ {
		if y+dy[i] >= 0 && y+dy[i] < n && x+dx[i] >= 0 && x+dx[i] < m && check[y+dy[i]][x+dx[i]] == 0 && data[y+dy[i]][x+dx[i]] == 1 {
			check[y+dy[i]][x+dx[i]] = 1
			dfs(y+dy[i], x+dx[i])
		}
	}
}

func main() {
	var t, k int
	ans := 0
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		ans = 0
		clear()
		fmt.Scanf("%d%d%d", &m, &n, &k)
		for j := 0; j < k; j++ {
			var y, x int
			fmt.Scanf("%d%d", &x, &y)
			data[y][x] = 1
		}
		for l := 0; l < n; l++ {
			for o := 0; o < m; o++ {
				if data[l][o] == 1 && check[l][o] == 0 {
					check[l][o] = 1
					dfs(l, o)
					ans++
				}
			}
		}
		fmt.Println(ans)
	}

}
