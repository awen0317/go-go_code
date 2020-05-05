package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
	"golang.org/x/net/context"
	"laonanhai/PhaseFive/day11/logagent/kafka"
)

var (
	tailObj *tail.Tail
)

type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	//为了能够退出t.run
	ctx context.Context
	cancelFun context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx,cancel :=	context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:  path,
		topic: topic,
		ctx:ctx,
		cancelFun:cancel,
	}
	tailObj.init() // 根据路径打开对应的路径信息
	return
}
func (t *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	//当gorouting执行函数推出的时候，goroutine就推出了
	go t.run()
}

// func Init(fileName string) (err error) {
// 	config := tail.Config{
// 		ReOpen:    true,                                 // 重新打开
// 		Follow:    true,                                 // 是否跟随
// 		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
// 		MustExist: false,                                // 文件不存在不报错
// 		Poll:      true,
// 	}
// 	tailObj, err = tail.TailFile(fileName, config)
// 	if err != nil {
// 		fmt.Println("tail file failed, err:", err)
// 		return
// 	}
// 	return
// }
// func ReadChan() <-chan *tail.Line {
// 	return tailObj.Lines
// }
// func (t *TailTask) ReadChan() <-chan *tail.Line {
// 	return tailObj.Lines
// }
func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Println("tail task:%s_%s 结束了。。。\n",t.path,t.topic)
			return
		case line := <-t.instance.Lines:
			// 3.3发往kafka
			// kafka.SendTiKafka(t.topic,line.Text)//函数调用函数
			//先把书局放到一个一个通道中
			kafka.SendToChan(t.topic,line.Text)

		}
	}
}
