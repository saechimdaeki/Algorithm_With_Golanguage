package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string
	fmt.Scanln(&s)
	sort.Sort(sort.StringsAreSorted(s))
	fmt.Println(s)
}
