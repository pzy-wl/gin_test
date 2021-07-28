package utils

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"github.com/mzky/zip"
	"io"
	"os"
	"path/filepath"
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

func IsZip(zipPath string) bool {
	f, err := os.Open(zipPath)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}

	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

// password值可以为空""
func Zip(zipPath, password string, fileList []string) error {
	if len(fileList) < 1 {
		return fmt.Errorf("将要压缩的文件列表不能为空")
	}
	fz, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(fz)
	defer zw.Close()

	for _, fileName := range fileList {
		fr, err := os.Open(fileName)
		if err != nil {
			return err
		}

		// 写入文件的头信息
		var w io.Writer
		if password != "" {
			w, err = zw.Encrypt(fileName, password, zip.AES256Encryption)
		} else {
			w, err = zw.Create(fileName)
		}

		if err != nil {
			return err
		}

		// 写入文件内容
		_, err = io.Copy(w, fr)
		if err != nil {
			return err
		}
	}
	return zw.Flush()
}

// password值可以为空""
// 当decompressPath值为"./"时，解压到相对路径
func UnZip(zipPath, password, decompressPath string) error {
	if !IsZip(zipPath) {
		return fmt.Errorf("压缩文件格式不正确或已损坏")
	}
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if password != "" {
			if f.IsEncrypted() {
				f.SetPassword(password)
			} else {
				return errors.New("must be encrypted")
			}
		}
		fp := filepath.Join(decompressPath, f.Name)
		dir, _ := filepath.Split(fp)
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}

		w, err := os.Create(fp)
		if nil != err {
			return err
		}

		fr, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(w, fr)
		if err != nil {
			return err
		}
		w.Close()
	}
	return nil
}
