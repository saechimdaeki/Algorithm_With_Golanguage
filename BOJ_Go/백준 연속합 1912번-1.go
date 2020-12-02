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
   arr:=[]int{0}
   dp:=make([]int,n)
   for i:=0; i<n; i++{
       var tmp int
       fmt.Fscan(bufin,&tmp)
       arr=append(arr,tmp)
   }
   for i:=0; i<n; i++{
       dp[i]=arr[i+1]
       if i==0 {
           continue
       }
       dp[i]=max(dp[i-1]+arr[i+1],dp[i])
   }
   var answer=dp[0]
   for i:=1; i<n; i++{
       if answer<dp[i]{
           answer=dp[i]
       }
   }
   fmt.Fprint(bufout,answer)
}