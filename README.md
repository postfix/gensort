# gensort
gensort generates Go sort function for T[]

[Sample usage](https://github.com/mozu0/gensort/blob/master/sample/a.go):

```
//go:generate gensort --output sort.int64.go --type int64 --package main --name sortInt64 --less "x > y" --eq "x == y"
```

[Sample of generated sort function](https://github.com/mozu0/gensort/blob/master/sample/sort.int64.go)
