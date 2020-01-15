package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var s string
var n, tmp int

func main() {
	defer bufout.Flush()
	fmt.Fscan(bufin, &s)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0':
			tmp = 0
			break
		case '1':
			tmp = 1
			break
		case '2':
			tmp = 2
			break
		case '3':
			tmp = 3
			break
		case '4':
			tmp = 4
			break
		case '5':
			tmp = 5
			break
		case '6':
			tmp = 6
			break
		case '7':
			tmp = 7
			break
		case '8':
			tmp = 8
			break
		case '9':
			tmp = 9
			break
		case 'A':
			tmp = 10
			break
		case 'B':
			tmp = 11
			break
		case 'C':
			tmp = 12
			break
		case 'D':
			tmp = 13
			break
		case 'E':
			tmp = 14
			break
		case 'F':
			tmp = 15
			break

		}
		for j := i; j < len(s)-1; j++ {
			tmp *= 16
		}
		n += tmp
	}
	fmt.Fprint(bufout, n)
}

//그냥%x로 출력하면 16진수다....
