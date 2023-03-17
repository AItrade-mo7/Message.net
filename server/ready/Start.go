package ready

import (
	"time"

	"Message.net/server/disposeTask"
	"Message.net/server/global"
	"github.com/EasyGolang/goTools/mCycle"
)

// 在这里 启动一个子进程，来进行目录的变化监听
func Start() {
	StartEmail()

	disposeTask.Treatment()
	mCycle.New(mCycle.Opt{
		Func:      CycleFunc,
		SleepTime: time.Minute * 10, // 20 分钟额外执行一次邮件同步
	}).Start()

	// 启动进任务 进程 监听，监听一次 接口的保存结果
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
	SyncEmailUseCount() // 结束了同步一次
}
