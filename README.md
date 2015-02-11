# gensort
gensort generates Go sort function for T[]

gensort can generate parallel quicksort function for slices of any types.
You can specify comparing function by --eq and --less flag.
For example, when you want to sort ints in decreasing order, you can use --less "x > y".
For large (> 1e6) slices, the generated sorting function typically runs 2 times faster (even without any parallelism) than the sort function in standard Go library.

Please refer to http://blog.golang.org/generate for usage of go generate command.

[Sample usage](https://github.com/mozu0/gensort/blob/master/sample/a.go):

```go
package main

// Generate a function that sorts int64 in decreasing order.
//go:generate gensort --output sort.int64.go --type int64 --package main --name sortInt64 --less "x > y" --eq "x == y"

import "fmt"

func main() {
	xs := []int64{3, 1, 2}
	sortInt64(xs)
	fmt.Println(xs)
}
```

[Sample of generated sort function](https://github.com/mozu0/gensort/blob/master/sample/sort.int64.go)
