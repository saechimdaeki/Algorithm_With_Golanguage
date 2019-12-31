package main

/*
boj stack 문제
input:
14
push 1
push 2
top
size
empty
pop
pop
pop
size
empty
pop
push 3
empty
top

output:
2
2
0
2
1
-1
0
1
-1
0
3
*/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack struct {
	head *Element
	size int
}
type Element struct {
	next  *Element
	value int
}

func (s *Stack) Push(x int) {
	e := Element{s.head, x}
	s.head = &e
	s.size++
}
func (s *Stack) Pop() int {
	if s.size == 0 {
		return -1
	} else {
		e := s.head
		s.head = e.next
		s.size--
		return e.value
	}
}
func (s *Stack) Size() int {
	return s.size
}
func (s *Stack) Empty() int {
	if s.size > 0 {
		return 0
	} else {
		return 1
	}
}
func (s *Stack) Top() int {
	if s.size == 0 {
		return -1
	} else {
		e := s.head
		return e.value
	}
}
func main() {
	var n int
	fmt.Scanf("%d", &n)
	var stack Stack
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	for i := 0; i < n; i++ {
		scanner.Scan()
		str := scanner.Text()
		switch str {
		case "push":
			scanner.Scan()
			v, _ := strconv.Atoi(scanner.Text())
			stack.Push(v)
		case "pop":
			fmt.Fprintf(w, "%d\n", stack.Pop())
		case "size":
			fmt.Fprintf(w, "%d\n", stack.Size())
		case "empty":
			fmt.Fprintf(w, "%d\n", stack.Empty())
		case "top":
			fmt.Fprintf(w, "%d\n", stack.Top())
		}
	}
	w.Flush()
}

/*
golang study용 이코드는 boj에 제출하지않음
*/
