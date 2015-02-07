package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

//go:generate gensort --name quickSort --output generated.sort.go

func TestSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	xs := randomArray(1e7)
	xs1 := append([]int(nil), xs...)
	xs2 := append([]int(nil), xs...)
	quickSort(xs1)
	sort.Ints(xs2)
	if !reflect.DeepEqual(xs1, xs2) {
		t.Error("Sort failed.")
	}
}

func randomArray(n int) (xs []int) {
	for i := 0; i < n; i++ {
		xs = append(xs, rand.Intn(1e6))
	}
	return
}
