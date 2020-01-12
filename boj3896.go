package main

import (
	"bufio"
	"fmt"
	"os"
)

var arr [1299710]int
var t, n, tmp, rst int
var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

func sosoo() {
	arr[0] = -1
	arr[1] = -1
	for i := 2; i < 1299710; i++ {
		arr[i] = i
	}
	for i := 2; i*i < 1299710; i++ {
		if arr[i] == i {
			for j := (i * i); j < 1299710; j += i {
				if arr[j] == j {
					arr[j] = i
				}
			}
		}
	}
}

func main() {
	defer bufout.Flush()
	fmt.Fscanf(bufin, "%d", &t)
	sosoo()
	for i := 0; i < t; i++ {
		fmt.Fscan(bufin, &n)
		if arr[n] == n {
			fmt.Fprintln(bufout, 0)
		} else {
			rst = 1
			tmp = n + 1
			for {
				if arr[tmp] == tmp {
					break
				}
				tmp++
				rst++
			}
			tmp = n - 1
			for {
				if arr[tmp] == tmp {
					break
				}
				tmp--
				rst++
			}
			fmt.Fprintln(bufout, rst+1)
		}
	}
}
