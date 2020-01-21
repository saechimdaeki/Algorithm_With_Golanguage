package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

var n, x, sum int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(bufin, &arr[i])
	}
	fmt.Fscan(bufin, &x)
	sort.Ints(arr)
	for i := 0; i < n; i++ {
		left := 0
		right := n - 1
		for {
			if left > right {
				break
			}
			mid := (left + right) / 2
			if arr[mid] == x-arr[i] {
				sum++
				break
			} else if arr[mid] < x-arr[i] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	sum /= 2
	fmt.Fprint(bufout, sum)
}
