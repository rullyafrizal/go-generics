package main

import (
	"log"
	"testing"
)

// any adalah alias baru dari interface{} kosong
// any baru bisa dipakai di go1.18
func Length[T any](arg T) T {
	log.Println(arg)

	return arg
}

func TestLength(t *testing.T) {
	var x int = 10
	Length[int](x)

	var y string = "hello"
	Length[string](y)
}
