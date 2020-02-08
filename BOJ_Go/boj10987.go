package main

import (
	"bufio"
	"fmt"
	"os"
)

var s string
var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var cnt int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &s)
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			cnt++
		}
	}
	fmt.Fprint(bufout, cnt)
}
