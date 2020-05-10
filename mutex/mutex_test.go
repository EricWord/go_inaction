package mutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
   @author cuiguangsong
   @email cuiguangsong@duxiaoman.com
   @department SIOD DBA GROUP
   @date  2020/5/10 下午19:08
   @since Go1.13.8
*/

type data struct {
	sync.Mutex
}

func (d *data) test(s string) {
	d.Lock()
	defer d.Unlock()
	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}

}

func TestMutex(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	var d data

	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}

func TestMultiMutex(t *testing.T) {
	var m sync.Mutex
	m.Lock()
	{
		m.Lock()
		m.Unlock()
	}
	m.Unlock()

}
