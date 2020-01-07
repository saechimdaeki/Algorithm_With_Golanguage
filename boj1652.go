package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var zari [100]string
var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

func solution1() int {
	result := 0
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if zari[i][j] == '.' {
				cnt++
			} else if zari[i][j] == 'X' && cnt >= 2 {
				result++
				cnt = 0
			} else if zari[i][j] == 'X' {
				cnt = 0
			}

		}
		if cnt >= 2 {
			result++
		}
	}
	return result

}
func solution2() int {
	result := 0
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if zari[j][i] == '.' {
				cnt++
			} else if zari[j][i] == 'X' && cnt >= 2 {
				result++
				cnt = 0
			} else if zari[j][i] == 'X' {
				cnt = 0
			}
		}
		if cnt >= 2 {
			result++
		}
	}
	return result
}
func main() {
	defer bufout.Flush()

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(bufin, &zari[i])
	}
	fmt.Fprintf(bufout, "%d %d", solution1(), solution2())
}
