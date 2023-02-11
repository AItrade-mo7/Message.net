package ready

import (
	"fmt"
	"time"

	"Message.net/server/global/config"
	"github.com/fsnotify/fsnotify"
)

var count = 0

// 监听目录变动

func ListenFolder() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(fmt.Errorf("ready.ListenFolder1 %v", err))
	}
	err = watcher.Add(config.Dir.TaskQueue)
	if err != nil {
		panic(fmt.Errorf("ready.ListenFolder2 %v", err))
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			count++
			fmt.Println(count, event, ok)
		case err := <-watcher.Errors:
			fmt.Println(222, err)
		case <-time.After(60 * time.Second):
			fmt.Println(3333, "time.After")
			continue
		}
	}
}

/*

if err != nil {
		log.Println("err11", err)
	}
	defer watcher.Close()

	done2 := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				fmt.Println("***********event")
				if !ok {
					return
				}
				fmt.Println("event.Op=>", event.Op)

				fmt.Println("文件操作类型判断是不是新建一个文件：", event.Op&fsnotify.Create == fsnotify.Create)
				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Println("*Create**event")
					fmt.Println("新的文件:", event.Name)
					mate := strings.Split(event.Name, "\\")
					fileName := mate[len(mate)-1]

					fmt.Println(1111, fileName)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			case <-time.After(60 * time.Second):
				continue
			}
		}
	}()
	err = watcher.Add(config.Dir.TaskQueue)
	if err != nil {
		log.Println("err22", err)
	}
	<-done2



*/
