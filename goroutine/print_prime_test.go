package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

/*
   @author eric
   @email 2635863525@qq.com
   @department SIOD DBA GROUP
   @date  2020/5/5 上午08:22
   @since Go1.13.8
*/

//wg用来等待程序完成
var wg sync.WaitGroup

//显示5000以内的素数值
func printPrime(prefix string) {

	//在函数退出时调用Done来通知调用函数工作已经完成
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}

		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Println("Completed", prefix)
}
func TestPrintPrime(t *testing.T) {

	//分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	//计数+2表示要等待两个goroutine
	wg.Add(2)
	//创建两个goroutine
	fmt.Println("Create goroutines")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("Terminating Program")

}
