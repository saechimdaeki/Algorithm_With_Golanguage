package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var s string
var check bool

func main() {
	defer bufout.Flush()
	check = true
	all := [8]string{"000", "001", "010", "011", "100", "101", "110", "111"}
	fmt.Fscan(bufin, &s)
	if len(s) == 1 && s[0]-'0' == 0 {
		fmt.Fprint(bufout, "0")
	}
	for i := 0; i < len(s); i++ {
		n := s[i] - '0'
		if check == true && n < 4 {
			if n == 0 {
				continue
			} else if n == 1 {
				fmt.Fprint(bufout, "1")
			} else if n == 2 {
				fmt.Fprint(bufout, "10")
			} else if n == 3 {
				fmt.Fprint(bufout, "11")
			}
			check = false
		} else {
			fmt.Fprint(bufout, all[n])
			check = false
		}
	}
}
