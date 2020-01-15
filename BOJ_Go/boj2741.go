package main

//// 그냥 fmt쓸시에 시간초과가 나므로 bufio사용해야함
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	var n int
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		w.WriteString(strconv.Itoa(i) + "\n")
	}
	w.Flush()
}
