package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

/*
   @author eric
   @email 2635863525@qq.com
   @department SIOD DBA GROUP
   @date  2020/5/10 下午15:43
   @since Go1.13.8
*/

type receiver struct {
	sync.WaitGroup
	data chan int
}

//工厂方法
func newReceiver() *receiver {
	r := &receiver{
		data: make(chan int),
	}
	r.Add(1)
	go func() {
		defer r.Done()
		//	接收消息，直到通道被关闭
		for x := range r.data {
			fmt.Println("receive:", x)

		}

	}()

	return r
}




func TestReceiver(t *testing.T) {
	r := newReceiver()
	r.data <- 1
	r.data <- 2
	//关闭通道，发出结束通知
	close(r.data)
	r.Wait()

}
