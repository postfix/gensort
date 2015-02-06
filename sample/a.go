package main

// Generate a function that sorts int64 in decreasing order.
//go:generate gensort --output sort.int64.go --type int64 --package main --name sortInt64 --less "x > y" --eq "x == y"

import "fmt"

func main() {
	xs := []int64{3, 1, 2}
	sortInt64(xs)
	fmt.Println(xs)
}
