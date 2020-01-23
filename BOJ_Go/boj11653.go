package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var arr [10000000]int

var n int

func main() {
	var vector []int
	defer bufout.Flush()
	fmt.Fscan(bufin, &n)
	arr[0] = -1
	arr[1] = -1
	for i := 2; i < 10000000; i++ {
		arr[i] = i
	}
	for i := 2; i*i < 10000000; i++ {
		if arr[i] == i {
			for j := i * i; j < 10000000; j += i {
				if arr[j] == j {
					arr[j] = i
				}
			}
		}
	}
	for i := 2; i < 10000000; i++ {
		if arr[i] == i {
			vector = append(vector, i)
		}
	}
	for i := 0; i < len(vector); i++ {
		for {
			if n%vector[i] != 0 {
				break
			}
			fmt.Fprintf(bufout, "%d\n", vector[i])
			n /= vector[i]
			if n == 1 {
				break
			}
		}
		if n == 1 {
			break
		}
	}
}
