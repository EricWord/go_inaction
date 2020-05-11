package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

/*
   @author eric
   @email 2635863525@qq.com
   @department SIOD DBA GROUP
   @date  2020/5/10 下午16:09
   @since Go1.13.8
*/
//用通道实现信号量

func TestSemaphore(t *testing.T) {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	//最多允许2个并发同时执行
	sem := make(chan struct{}, 2)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			//acquire:获取信号
			sem <- struct{}{}
			//release:释放信号
			defer func() {
				<-sem
			}()

			time.Sleep(time.Second * 2)
			fmt.Println(id, time.Now())

		}(i)

	}

	wg.Wait()
}
