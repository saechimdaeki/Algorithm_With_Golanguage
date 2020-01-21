package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n, res int
var arr [50][50]rune

func counting(y, x int, ck bool) {
	count := 1
	for i := 0; i < n-1; i++ {
		if ck {
			if arr[i][x] == arr[i+1][x] {
				count++
				if res < count {
					res = count
				}
			} else {
				count = 1
			}
		} else {
			if arr[y][i] == arr[y][i+1] {
				count++
				if res < count {
					res = count
				}
			} else {
				count = 1
			}
		}
	}
}

func check(y1, x1, y2, x2 int) {
	arr[y1][x1], arr[y2][x2] = arr[y2][x2], arr[y1][x1]
	isYmove := true
	if x1 != x2 {
		isYmove = false
	}
	if isYmove {
		counting(y1, x1, false)
		counting(y2, x1, false)
		counting(y1, x1, true)
	} else {
		counting(y1, x1, true)
		counting(y1, x2, true)
		counting(y1, x1, false)
	}
	arr[y1][x1], arr[y2][x2] = arr[y2][x2], arr[y1][x1]
}
func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	var str string
	for i := 0; i < n; i++ {
		fmt.Fscan(bufin, &str)
		for j := 0; j < n; j++ {
			arr[i][j] = rune(str[j])
		}
	}

	for y := 0; y < n; y++ {
		for x := 0; x < n-1; x++ {
			check(y, x, y, x+1)
			check(x, y, x+1, y)
		}
	}

	fmt.Fprint(bufout, res)
}
