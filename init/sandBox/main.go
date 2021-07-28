package main
import "fmt"

//同一个包不同源文件的init函数执行顺序，golang spec没做说明，以上述程序输出来看，执行顺序是源文件名称的字典序。

var _ int64 = s()
func init() {
	fmt.Println("init in sandbox.go")
}
func s() int64 {
	fmt.Println("calling s() in sandbox.go")
	return 1
}
func main() {
	fmt.Println("main")
}