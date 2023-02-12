package ready

import (
	"fmt"

	"Message.net/server/disposeTask"
	"Message.net/server/global"
)

// 在这里 启动一个子进程，来进行目录的变化监听
func Start() {
	fmt.Println("ready.Start")

	// 读取一次目录的任务列表
	disposeTask.Treatment()

	go WatchTaskDir()
}

func WatchTaskDir() {
	for {
		TaskID, ok := <-global.TaskChan
		if !ok {
			break
		}
		fmt.Println("新任务进来了", TaskID)
		disposeTask.Treatment()
	}
}
