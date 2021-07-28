package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeToUnix(t *testing.T) {
	//时间和时间戳相互转换
	num := time.Now().Unix()
	fmt.Println("time is:", num)
	tm := time.Unix(1615865410, 0)
	fmt.Println(tm.Format("2006-01-02 15:04:05"))
}
