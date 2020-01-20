package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

var a,b int64
func gcd(x,y int64)int64 {
	if y==0 {
		return x
	}else {
		return gcd(y,x%y)
	}
}



func main(){
	defer bufout.Flush()
	
	fmt.Fscan(bufin,&a,&b)
	fmt.Fprint(bufout,(a*b)/gcd(a,b))
	
}