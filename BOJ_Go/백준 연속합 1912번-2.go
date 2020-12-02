package main

import (
	"bufio"
	"fmt"
	"os"
)
var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var n int
func max(x,y int) int{
    if x<y{
        return y
    }
    return x
}
func main() {
   defer bufout.Flush()
   fmt.Fscan(bufin,&n)
   arr:=make([]int,n)
   dp:=make([]int,n)
   for i:=0; i<n; i++{
       fmt.Fscan(bufin,&arr[i])
   }
   for i:=0; i<n; i++{
       dp[i]=arr[i]
       if i==0 {
           continue
       }
       dp[i]=max(dp[i-1]+arr[i],dp[i])
   }
   var answer=dp[0]
   for i:=1; i<n; i++{
       if answer<dp[i]{
           answer=dp[i]
       }
   }
   fmt.Fprint(bufout,answer)
}