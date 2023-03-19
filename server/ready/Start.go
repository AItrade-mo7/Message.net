package ready

import (
	"time"

	"Message.net/server/disposeTask"
	"Message.net/server/global"
	"github.com/EasyGolang/goTools/mCycle"
)

// 在这里 启动一个子进程，来进行目录的变化监听
func Start() {
	go WatchTaskDir() // 启动监听进程

	mCycle.New(mCycle.Opt{
		Func:      CycleFunc,
		SleepTime: time.Minute * 10, // 20 分钟额外执行一次邮件同步
	}).Start()

	// 发送一次邮件
	StartEmail()
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

func CycleFunc() {
	SyncEmailUseCount() // 结束了同步一次
}
