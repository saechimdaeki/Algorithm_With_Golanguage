package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n, m, k, cnt, sum int
var check bool

func main() {
	defer bufout.Flush()
	fmt.Fscanf(bufin, "%d", &n)
	arr := make([]int, n+1)
	fmt.Fscan(bufin, &m, &k)

	for i := 0; i < n+1; i++ {
		fmt.Fscanf(bufin, "%d", &arr[i])
	}

	sort.Ints(arr)
	for i := n; i > 0; i-- {

		sum += arr[i]
		cnt++
		if sum >= (m * k) {
			check = true
			break
		}
	}
	if check == true {
		fmt.Fprint(bufout, cnt)
	} else {
		fmt.Fprint(bufout, "STRESS")
	}

}
