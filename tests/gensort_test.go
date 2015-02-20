package tests

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

//go:generate gensort --name quickSort --output generated.sort.go --package=tests

//go:generate gensort --name cmpQuickSort --cmp "x - y" --output generated.cmp.sort.go --package=tests

func TestSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	xs := randomArray(1e6)

	{ // quickSort
		xs1 := append([]int(nil), xs...)
		xs2 := append([]int(nil), xs...)
		quickSort(xs1)
		sort.Ints(xs2)
		if !reflect.DeepEqual(xs1, xs2) {
			t.Error("Sort failed.")
		}
	}

	{ // cmpQuickSort
		xs1 := append([]int(nil), xs...)
		xs2 := append([]int(nil), xs...)
		cmpQuickSort(xs1)
		sort.Ints(xs2)
		if !reflect.DeepEqual(xs1, xs2) {
			t.Error("Sort failed.")
		}
	}
}

func randomArray(n int) (xs []int) {
	for i := 0; i < n; i++ {
		xs = append(xs, rand.Intn(1e5))
	}
	return
}
