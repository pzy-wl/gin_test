package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestByline(t *testing.T) {
	f, err := os.OpenFile("C:\\Users\\Administrator\\Desktop\\123.txt", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open f error!", err)
		return
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')

		fmt.Println("line1", line)
		line = strings.TrimSpace(line)
		fmt.Println("line2", line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read f error!")
				return
			}
		}
	}
}
