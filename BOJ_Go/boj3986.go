package main

import (
	"bufio"
	"fmt"
	"os"
)
/*
	스택문제 go로풀기 너무어려워요 ㅠㅠㅠㅠㅠ 
*/
var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

var n,cnt int 

func main(){
	defer bufout.Flush()
	fmt.Fscan(bufin,&n)
	for i:=0; i<n; i++ {
		var s string 
		fmt.Fscan(bufin,&s)
		var stack[] byte 
		
		
		for j:=0; j<len(s); j++ {
			n:=len(stack)-1
			if len(stack)>0 && stack[n]==s[j]{
				stack=stack[:n]
			}else{
				stack=append(stack,s[j])
			}
		}
		if len(stack)==0{
			cnt++
		}
	}
	fmt.Fprint(bufout,cnt)
}