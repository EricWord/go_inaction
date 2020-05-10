package goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

/*
   @author cuiguangsong
   @email cuiguangsong@duxiaoman.com
   @department SIOD DBA GROUP
   @date  2020/5/5 上午08:55
   @since Go1.13.8
*/

var (
	//所有goroutine都要增加其值的变量
	counter int
	//用来等待程序结束
)

//增加counter变量的值
func inCounter(id int) {
	//函数退出时调用Done来通知调用函数工作已经完成
	defer wg.Done()
	for count := 0; count < 2; count++ {
		//捕获counter的值
		value := counter
		//当前goroutine从线程退出，并放回到队列
		runtime.Gosched()
		//增加本地变量value的值
		value++

		//将该值保存回counter
		counter = value
	}

}
func TestRace(t *testing.T) {
	wg.Add(2)
	//创建两个goroutine
	go inCounter(1)
	go inCounter(2)
	//等待goroutine结束
	wg.Wait()
	fmt.Println("Final counter:", counter)

}
