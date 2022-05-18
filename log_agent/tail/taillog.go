package tail

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"log_agent_ty/kafka"
)

var (
	tailObj *tail.Tail
	err     error
)

//one catch log task instance
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	//为了控制t.run()退出
	ctx       context.Context
	cancelFun context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {

	fmt.Println("new task : ", path)
	ctx, cancel := context.WithCancel(context.Background())

	tailObj = &TailTask{
		path:      path,
		topic:     topic,
		ctx:       ctx,
		cancelFun: cancel,
	}

	//start reading file
	tailObj.init()
	return
}

func (t *TailTask) init() {
	//log read config
	config := tail.Config{
		Location:    &tail.SeekInfo{Offset: 0, Whence: 0}, //从文件的某个地方开始读
		ReOpen:      true,                                 //重新打开 Reopen recreated files (tail -F)
		MustExist:   false,                                //文件不存在报错 Fail early if the file does not exist
		Poll:        false,
		Pipe:        false,
		RateLimiter: nil,
		Follow:      true, //是否跟随
		MaxLineSize: 0,
		Logger:      nil,
	}

	//listen file write
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
	}
	//当goroutine 执行的函数退出时,goroutine结束
	go t.run() //read file
}

//read log and send
func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Println("tail task finished :", t.path + t.topic)
			return
		case line := <-t.instance.Lines:
			//use chan send msg
			kafka.SendToChan(t.topic, line.Text)
		}
	}
}
