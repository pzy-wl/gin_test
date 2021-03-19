package main

//定时任务
import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {
	//定时任务
	c := cron.New()
	spec := "@every 1s" // 每1h执行
	c.AddFunc(spec, func() {
		//TODO 等到控制台好了打开测试
		//handlePeerStatus()
		fmt.Println("心跳包发送成功!")
	})
	c.Start()
	select {}
}
