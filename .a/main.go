package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"
)

import "C"

//#cgo LDFLAGS: -L${SRCDIR} -L /opt/sgxsdk/lib64 -lwrapper -l sgx_urts -ldl
//#include <stdint.h>
//extern int32_t rust_sgx_mpc(char* c, size_t l);
//extern int32_t rust_sgx_rsakey(char* c, size_t l, char* d);
//extern unsigned long long init_enclave();
func main() {
	a := make([]string, 0)
	n := make([]int, 0)
	b1, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化字符数组时出错%v", err.Error())
		return
	}
	b2, err := json.Marshal(n)
	if err != nil {
		fmt.Printf("序列化数组时出错%v", err.Error())
		return
	}
	argsString1 := string(b1)
	argsString2 := string(b2)
	p1 := (*reflect.StringHeader)(unsafe.Pointer(&argsString1))
	p2 := (*reflect.SliceHeader)(unsafe.Pointer(&argsString2))
	res, _ := C.rust_sgx_mpc((*C.char)(unsafe.Pointer(p1.Data)), C.ulong(len(argsString1)), (*C.char)(unsafe.Pointer(p2.Data)))
	fmt.Println("执行结果是：%v", res)
}
