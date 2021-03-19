package mashall

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	//测试测试unmarshal
	name := "hello world"
	n, err := json.Marshal(name)
	if err != nil {
		fmt.Println("this is an error")
		return
	}
	l := make([]string, 0)
	json.Unmarshal(n, l)
	fmt.Println("切片是：", l)
	for _, v := range l {
		fmt.Println(v)
	}
}
