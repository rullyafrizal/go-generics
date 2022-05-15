package main

import (
	"log"
	"testing"
)

func IsSame[T comparable](a, b T) string {
	if a == b {
		return "sama"
	}

	return "beda"
}

// any bukan comparable, jadi tidak bisa menggunakan operator perbandingan
// func IsSame2[T any](a, b T) string {
// 	if a == b {
// 		return "sama"
// 	}

// 	return "beda"
// }

func TestIsSame(t *testing.T) {
	log.Println(IsSame[int](10, 10))
	log.Println(IsSame[string]("hello", "hello"))
	log.Println(IsSame[int](10, 20))
	log.Println(IsSame[string]("hello", "world"))
}
