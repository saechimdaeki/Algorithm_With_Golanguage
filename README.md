# study_Go
- for study go language  
- go lang 심화튜토리얼 :http://golang.site/
- golang tutorial study link : https://go-tour-kr.appspot.com/
- go lang ver.

![go](./go.JPG)

## go언어를 실행 할때는 터미널에서 go run 파일명.go

### 왜 go언어를 해야하는가 ? 
- https://blog.jse.li/posts/torrent/
- https://blog.youngbin.xyz/2019-09-09-migrating-skhus-backend-from-nodejs-to-golang/
- https://github.com/golang/go/wiki/LearnServerProgramming
- https://www.buzzvil.com/ko/2018/02/12/tech-blog-go-%EC%84%9C%EB%B2%84-%EA%B0%9C%EB%B0%9C%ED%95%98%EA%B8%B0/

## go언어의 간단한 실용예
- http://golang.site/go/article/111-%EA%B0%84%EB%8B%A8%ED%95%9C-%EC%9B%B9-%EC%84%9C%EB%B2%84-HTTP-%EC%84%9C%EB%B2%84 (서버)
- http://pyrasis.com/book/GoForTheReallyImpatient/Unit58   (서버)
- https://mingrammer.com/building-blockchain-in-go-part-1/ (golang으로 blockchain만들기)

## golang for backend
- Golang의 독특한 특징 3가지 - A declared are not used, Multiple return values, Timeout (https://voidmainvoid.tistory.com/128)

## go언어로 bfs 문제를 푸는 예시 (직접작성)

### go언어로 백준알고리즘 문제풀기 [BOJGolang](./BOJ_Go)

- 문제:

![ex](./ex.JPG)

<code>
    
    package main
        import (
	    "bufio"
	    "fmt"
	    "os"
        )

        var n, m, num1, num2, x, y int
        var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
        var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
        var grid [101][101]int
        var visited [101]bool
        var count [101]int

        func bfs(start int) {
            queue := make([]int, 0)
            queue = append(queue, start)

            for len(queue) > 0 {
                cur := queue[0]
                queue = queue[1:]
            for i := 1; i <= n; i++ {
                if grid[cur][i] == 1 && visited[i] == false {
                    visited[i] = true
                    count[i] = count[cur] + 1
                    queue = append(queue, i)
                    }
                }   
            }
        }
        func main() {
            defer bufout.Flush()
            fmt.Fscan(bufin, &n)
            fmt.Fscan(bufin, &num1, &num2)
            fmt.Fscan(bufin, &m)
            var a1, a2 int
            for i := 1; i <= m; i++ {
            fmt.Fscan(bufin, &a1, &a2)
            grid[a1][a2] = 1
            grid[a2][a1] = 1
        }
            bfs(num1)
            if count[num2] != 0 {
                fmt.Fprintln(bufout, count[num2])
            } else {
                fmt.Fprintln(bufout, "-1")
            }
        }

</code>

- 위의 코드와 같이 bufio reader writer를통해 시간을 엄청나게 단축시킬수있음 단순 fmt.Print,fmt.Scanf를 써서 문제를 풀었을시 12ms. bufio를 사용하여 reader를통해 받는 fmt.fScanf를사용시 4ms.  go 언어로 알고리즘 로직이 맞지만 시간초과가날시에 이를통해 해결한 사례가 많음. 

![time](./time.JPG)


- for문은 그때그때 따라서 ~ 

<code>
	
	for _, number := range numbrs{ }
	
	for i:=0; i<len(numbers); i++ { }
	
</code>	

