package main

import (
	"fmt"
	"testing"
)

func MultipleParameter[T1 any, T2 any](arg1 T1, arg2 T2) {
	fmt.Println(arg1, arg2)
}

func TestMultipleParameter(t *testing.T) {
	MultipleParameter[int, string](10, "hello")
	MultipleParameter[[]string, []int](
		[]string{"hello", "world"},
		[]int{1, 2, 3},
	)
}
