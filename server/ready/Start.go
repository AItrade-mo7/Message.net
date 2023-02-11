package ready

import (
	"fmt"

	"Message.net/server/disposeTask"
)

// 在这里 启动一个子进程，来进行目录的变化监听
func Start() {
	fmt.Println("ready.Start")

	// 读取一次目录的任务列表
	disposeTask.Treatment()
}
