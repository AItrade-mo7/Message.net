package ready

import (
	"time"

	"Message.net/server/disposeTask"
	"Message.net/server/global"
	"github.com/EasyGolang/goTools/mCycle"
)

// 在这里 启动一个子进程，来进行目录的变化监听
func Start() {
	go StartEmail()

	mCycle.New(mCycle.Opt{
		Func:      CycleFunc,
		SleepTime: time.Minute * 10, // 10 分钟额外执行一次检查和同步
	}).Start()

	// 启动进任务进程监听
	go WatchTaskDir()
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
	SyncEmailUseCount()     // 邮件频率 同步
	disposeTask.Treatment() // Task 目录检查
}
