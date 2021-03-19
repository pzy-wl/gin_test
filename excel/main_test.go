package main

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	s := "123"
	s2 := "123"
	fmt.Println(s == s2)
}
