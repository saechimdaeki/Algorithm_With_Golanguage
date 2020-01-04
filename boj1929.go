package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func sosoo(x int) int {
	if x == 1 {
		return 0
	}

	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return 0
		}
	}
	return 1
}

func main() {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)
	w := bufio.NewWriter(os.Stdout)
	for i := a; i <= b; i++ {
		if sosoo(i) == 1 {
			fmt.Fprintf(w, "%d\n", i)
		}
	}
	w.Flush()

}
