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
var n int 


func main(){
	defer bufout.Flush()
	fmt.Fscan(bufin,&n)
	for i:=0; i<n; i++ {
	fmt.Fscan(bufin,&a,&b)
	fmt.Fprintln(bufout,(a*b)/gcd(a,b))
	
	}
}