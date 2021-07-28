// Copyright (c) 2016 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package main

import (
	"bytes"
	"crypto"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"testing"
)

// An example for the ECDH key-exchange using the curve P256.
func ExampleGeneric() {
	p256 := Generic(elliptic.P256())

	privateAlice, publicAlice, err := p256.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Failed to generate Alice's private/public key pair: %s\n", err)
	}
	privateBob, publicBob, err := p256.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Failed to generate Bob's private/public key pair: %s\n", err)
	}
	if err := p256.Check(publicBob); err != nil {
		fmt.Printf("Bob's public key is not on the curve: %s\n", err)
	}
	secretAlice := p256.ComputeSecret(privateAlice, publicBob)
	if err := p256.Check(publicAlice); err != nil {
		fmt.Printf("Alice's public key is not on the curve: %s\n", err)
	}
	secretBob := p256.ComputeSecret(privateBob, publicAlice)

	if !bytes.Equal(secretAlice, secretBob) {
		fmt.Printf("key exchange failed - secret X coordinates not equal\n")
	}
	fmt.Println("密钥协商结果是：", bytes.Equal(secretAlice, secretBob))
	// Output:密钥协商结果是： true
}

func BenchmarkP256(b *testing.B) {
	p256 := Generic(elliptic.P256())
	privateAlice, _, err := p256.GenerateKey(rand.Reader)
	if err != nil {
		b.Fatalf("Failed to generate Alice's private/public key pair: %s", err)
	}
	_, publicBob, err := p256.GenerateKey(rand.Reader)
	if err != nil {
		b.Fatalf("Failed to generate Bob's private/public key pair: %s", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p256.ComputeSecret(privateAlice, publicBob)
	}
}

func BenchmarkKeyGenerateP256(b *testing.B) {
	p256 := Generic(elliptic.P256())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, err := p256.GenerateKey(rand.Reader)
		if err != nil {
			b.Fatalf("Failed to generate Alice's private/public key pair: %s", err)
		}
	}
}

func TestECDH(t *testing.T) {
	p256 := Generic(elliptic.P256())

	privateAlice, publicAlice, err := p256.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Failed to generate Alice's private/public key pair: %s\n", err)
		return
	}
	privateBob, publicBob, err := p256.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Failed to generate Bob's private/public key pair: %s\n", err)
		return
	}
	a, err := json.Marshal(privateAlice)
	if err != nil {
		fmt.Printf("marshall err%v", err.Error())
		return
	}
	var tmpa crypto.PrivateKey
	err = json.Unmarshal(a, &tmpa)
	if err != nil {
		fmt.Printf("unmarshall err:%v", err.Error())
		return
	}
	fmt.Printf("类型是:%T, %v\n", tmpa, tmpa)
	b, err := json.Marshal(privateBob)
	if err != nil {
		fmt.Printf("marshall err%v", err.Error())
		return
	}
	fmt.Println("私钥比较结果是", bytes.Equal(a, b))
	fmt.Println("私钥是", string(a), string(b))
	if err := p256.Check(publicBob); err != nil {
		fmt.Printf("Bob's public key is not on the curve: %s\n", err)
		return
	}
	secretAlice := p256.ComputeSecret(&tmpa, publicBob)
	if err := p256.Check(publicAlice); err != nil {
		fmt.Printf("Alice's public key is not on the curve: %s\n", err)
		return
	}
	secretBob := p256.ComputeSecret(privateBob, publicAlice)
	if !bytes.Equal(secretAlice, secretBob) {
		fmt.Printf("key exchange failed - secret X coordinates not equal\n")
		return
	}
	fmt.Println("密钥协商结果是：", bytes.Equal(secretAlice, secretBob))
}
