package main
import "fmt"

//init函数不可以被调用，上面代码会提示：undefined: init

func init() {
	fmt.Println("init")
}
func main() {
	//init()
}
