package goroutine

import (
	"fmt"
	"os"
	"testing"
	"time"
)

/*
   @author cuiguangsong
   @email cuiguangsong@duxiaoman.com
   @department SIOD DBA GROUP
   @date  2020/5/10 下午16:15
   @since Go1.13.8
*/

func TestTick(t *testing.T) {
	go func() {
		for {
			select {

			case <-time.After(time.Second * 5):
				fmt.Println("timeout...")
				os.Exit(0)
			}
		}
	}()

	go func() {
		tick := time.Tick(time.Second)
		for {
			select {
			case <-tick:
				fmt.Println(time.Now())
			}
		}
	}()

	<-(chan struct{})(nil)
}
