package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

func min(x int ,y int)int {
	if x>y {
		return y
	}else{
		return x
	}
}

const infinite int = 2e9+1
var n,m,answer int 
var grid [50]string 
var white=[8]string{
"WBWBWBWB","BWBWBWBW",
"WBWBWBWB","BWBWBWBW",
"WBWBWBWB","BWBWBWBW",
"WBWBWBWB","BWBWBWBW",
}

var black=[8]string{
"BWBWBWBW","WBWBWBWB",
"BWBWBWBW","WBWBWBWB",
"BWBWBWBW","WBWBWBWB",
"BWBWBWBW","WBWBWBWB",
}



func checkwhite(x int ,y int )int {
	cnt:=0
	for i:=x; i<x+8; i++ {
		for j:=y; j<y+8; j++{
			if grid[i][j] != white[i-x][j-y]{
				cnt++
			}
		}
	}
	return cnt 
}


func checkblack(x int ,y int )int {
	cnt:=0
	for i:=x; i<x+8; i++ {
		for j:=y; j<y+8; j++{
			if grid[i][j] != black[i-x][j-y]{
				cnt++
			}
		}
	}
	return cnt 
}

func main(){
	defer bufout.Flush()
	fmt.Fscan(bufin,&n,&m)
	for i:=0; i<n; i++{
		fmt.Fscan(bufin,&grid[i])
	}
	answer=infinite
	for i:=0; i<=n-8; i++{
		for j:=0; j<=m-8; j++ {
		answer=	min(answer,min(checkblack(i,j),checkwhite(i,j)))
		}
	}
	fmt.Fprint(bufout,answer)
	
	
	
	//fmt.Fprintln(bufout,white[1])
	//fmt.Fprint(bufout,infinite)
}