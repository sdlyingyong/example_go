package tail

import (
	"fmt"
	"log_agent_ty/etcd"
)

var tskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntryConf,
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry),
	}

	for _, logEntry := range logEntryConf {
		//初始化有几个tailtask要记录,为了后续知道哪个是新增的
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		key := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.tskMap[key] = tailObj
	}

	//监听热更新配置
	go tskMgr.run()
}

//listen new conf chann ,if  new conf tell and execute
func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:

			for _, conf := range newConf {
				fmt.Println("热更新被传递进来了: ", conf)
				key := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := tskMgr.tskMap[key]
				if ok {
					//原来就有,不需要操作
					fmt.Println("原来就有,不需要操作")
					continue
				} else {
					//配置新增
					tskMgr.tskMap[key] = NewTailTask(conf.Path,conf.Topic)
					fmt.Println("配置新增了 :",key)
				}
			}

			//配置删除
			//拿到现有的任务,对比热更新的
			//如果热更新有,就不用删除
			//如果没出现在热更新中,就要删除
			for _,c1 := range t.logEntry{
				isDelete := true

				for _,c2 := range newConf {
					if c2.Path == c1.Path && c2.Topic == c1.Topic {
						isDelete = false
						continue
					}
				}

				if isDelete {
					key := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					t.tskMap[key].cancelFun()
					delete(t.tskMap,key)
				}
			}
		}
	}

}

func NewConfChan() chan []*etcd.LogEntry {
	return tskMgr.newConfChan
}
