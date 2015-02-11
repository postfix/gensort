# gensort
gensort generates Go sort function for T[]

Sample usage:

//go:generate gensort --output sort.int64.go --type int64 --package main --name sortInt64 --less "x > y" --eq "x == y"
