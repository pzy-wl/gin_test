package pack

import (
	"bee/init/test_util"
	"fmt"
)


//init 函数先执行被依赖的包中的init函数
var Pack int = 6

func init() {
	a := test_util.Util
	fmt.Println("init pack ", a)
}
