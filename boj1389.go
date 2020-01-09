package main

import (
	"bufio"
	"os"
	"fmt"
)

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n, m int
var grid [101][101]int

func solve() {
	for i := 1; i <= n; i++ {
		for j:=1; j<=n; j++ {
			for k:=1; k<=n; k++ {
				if j==k{
					continue
				} else if grid[j][i]>=1 && grid[i][k]>=1  {
					if grid[j][k]==0 {
						grid[j][k]=grid[j][i]+grid[i][k]
					} else {
						grid[j][k]=min(grid[j][k],grid[j][i]+grid[i][k])
					}
				}
			}
		}

	}
}
var a,b ,s1,s2 int
func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin,&n,&m)
	for i:=0; i<m; i++{
		fmt.Fscan(bufin,&a,&b)
		grid[a][b]=1
		grid[b][a]=1
	}
	solve()
	s1=987654321
	for i:=1; i<=n; i++ {
		sum:=0
		for j:=1; j<=n; j++ {
			sum+=grid[i][j]
		}
		if s1 > sum {
			s1=sum
			s2=i
		}
	}
	fmt.Fprintln(bufout,s2)
}
