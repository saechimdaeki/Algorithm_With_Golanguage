package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n int

/*
func solvejump(x int) int {
	if x == n-1 {
		return 0
	}
	if x >= n {
		return
	}
}
*/
const infinite int = 2e9 + 1

var dynamic [1001]int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	dynamic[1] = 1
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(bufin, &a)
		if dynamic[i] == 0 {
			break
		}
		for j := 1; j <= a; j++ {
			if i+j > 1000 {
				break
			}
			if dynamic[i+j] > dynamic[i]+1 || dynamic[i+j] == 0 {
				dynamic[i+j] = dynamic[i] + 1
			}

		}

	}
	if dynamic[n] == 0 {
		fmt.Fprint(bufout, "-1")
	} else {
		fmt.Fprint(bufout, dynamic[n]-1)
	}

}
