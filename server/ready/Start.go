package ready

import (
	"Message.net/server/disposeTask"
	"Message.net/server/global"
)

// 在这里 启动一个子进程，来进行目录的变化监听
func Start() {
	// 读取一次目录的任务列表
	disposeTask.Treatment()
	go WatchTaskDir()

	SyncEmailUseCount()
}

func WatchTaskDir() {
	for {
		TaskID, ok := <-global.TaskChan
		if !ok {
			break
		}
		global.Run.Println("=====新任务进来了======", TaskID)
		disposeTask.Treatment()
	}
}
