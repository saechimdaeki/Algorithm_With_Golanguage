package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer bufout.Flush()
	var line int
	fmt.Scanf("%d", &line)
	vector := make([]float64, line)
	var max float64
	var sum float64
	for i := 0; i < line; i++ {
		fmt.Fscan(bufin, &vector[i])
		if vector[i] > max {
			max = vector[i]
		}
	}
	for i := 0; i < line; i++ {
		sum += (vector[i] / max) * 100
	}
	fmt.Printf("%.2f", sum/float64(line))

}
