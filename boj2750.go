package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bufout.Flush()
	var n int
	fmt.Fscanln(bufin, &n)
	mySlice := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscanln(bufin, &a)
		mySlice = append(mySlice, a)
	}
	sort.Ints(mySlice)
	for _, i := range mySlice {
		fmt.Fprintln(bufout, i)
	}
}
