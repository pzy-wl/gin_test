package pack

import (
	"fmt"
	"ginTest/init/test_util"
)

//init 函数先执行被依赖的包中的init函数
var Pack int = 6

func init() {
	a := test_util.Util
	fmt.Println("init pack ", a)
}
