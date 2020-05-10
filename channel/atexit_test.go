package channel

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
)

/*
   @author cuiguangsong
   @email cuiguangsong@duxiaoman.com
   @department SIOD DBA GROUP
   @date  2020/5/10 下午16:22
   @since Go1.13.8
*/

//捕获INT、TERM信号，顺便实现一个简易的atexit函数

var exits = &struct {
	sync.RWMutex
	funcs   []func()
	signals chan os.Signal
}{}

func atexit(f func()) {
	exits.Lock()
	defer exits.Unlock()
	exits.funcs = append(exits.funcs, f)
}

func waitExit() {
	if exits.signals == nil {
		exits.signals = make(chan os.Signal)
		signal.Notify(exits.signals, syscall.SIGINT, syscall.SIGTERM)
	}

	exits.RLock()
	for _, f := range exits.funcs {
		//即便某些函数panic,延迟调用也能确保后续函数执行
		defer f()
	}

	exits.RUnlock()
	<-exits.signals

}

func TestAtexit(t *testing.T) {
	atexit(func() {
		println("exit1...")
	})

	atexit(func() {
		println("exit2...")
	})

	waitExit()

}
