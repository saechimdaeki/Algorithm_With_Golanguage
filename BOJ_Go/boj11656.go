package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

var s string

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &s)
	srr := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		srr[i] = string(s[i:])
	}
	sort.Strings(srr)
	for i := 0; i < len(s); i++ {
		fmt.Fprintf(bufout, "%s\n", srr[i])

	}
}
