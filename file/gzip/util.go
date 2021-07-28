package main

import (
	"bytes"
	"compress/gzip"
	"io"
)

//GZipBytes 压缩
func GZipBytes(data []byte) []byte {
	var input bytes.Buffer
	g := gzip.NewWriter(&input)
	g.Write(data)
	g.Close()
	return input.Bytes()
}

//UGZipBytes 解压
func UGZipBytes(data []byte) []byte {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write(data)
	r, _ := gzip.NewReader(&in)
	r.Close()
	io.Copy(&out, r)
	return out.Bytes()
}
