package main

import (
	"bytes"
	"github.com/luke-park/ecdh25519"
)

func main() {
	prv1, err := ecdh25519.GenerateKey()
	if err != nil {
		println(err.Error())
		return
	}
	prv2, err := ecdh25519.GenerateKey()
	if err != nil {
		println(err.Error())
		return
	}
	s1 := prv1.ComputeSecret(prv2.Public())
	s2 := prv2.ComputeSecret(prv1.Public())
	if bytes.Compare(s1, s2) == 0 {
		println("fail")
		return
	}
	println("success")
}
