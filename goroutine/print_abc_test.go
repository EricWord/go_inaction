package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

/*
   @author cuiguangsong
   @email cuiguangsong@duxiaoman.com
   @department SIOD DBA GROUP
   @date  2020/4/30 上午10:18
   @since Go1.13.8
*/

func TestGoroutine(t *testing.T) {
	//	分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(2)
	//wg用来等待程序执行完成
	var wg sync.WaitGroup
	//计数加2表示要等待两个程序执行完成
	wg.Add(2)
	fmt.Println("start goroutines")

	//声明一个匿名函数，并创建一个goroutine
	go func() {
		//在函数退出时调用done来通知main函数工作已经完成
		defer wg.Done()
		//显示字母3次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)

			}
		}
	}()

	//声明一个匿名函数，并创建一个goroutine
	go func() {
		//在函数退出时调用done来通知main函数工作已经完成
		defer wg.Done()
		//显示字母3次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)

			}
		}

	}()

	//等待goroutine结束
	fmt.Println("waiting to finish")
	wg.Wait()
	fmt.Println("\n Terminating Program")

}
