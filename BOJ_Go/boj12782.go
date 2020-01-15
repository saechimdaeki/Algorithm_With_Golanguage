package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n int
var s1, s2 string

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	for k := 0; k < n; k++ {
		fmt.Fscan(bufin, &s1, &s2)
		var number [1000000]int
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				number[s1[i]]++
			}
		}
		if number['0'] < number['1'] {
			fmt.Fprintf(bufout, "%d\n", number['1'])
		} else {
			fmt.Fprintf(bufout, "%d\n", number['0'])

		}
	}
}
