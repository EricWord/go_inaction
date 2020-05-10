package slice

import (
	"fmt"
	"testing"
	"time"
)

/*
   @author cuiguangsong
   @email cuiguangsong@duxiaoman.com
   @department SIOD DBA GROUP
   @date  2020/4/26 上午10:21
   @since Go1.13.8
*/

func TestSlice(t *testing.T) {
	//使用make创建空的整型切片
	slice1 := make([]int, 0)
	//使用切片字面量创建空的整型切片
	slice2 := []int{}

	fmt.Println(slice1)
	fmt.Println(slice2)

}

func TestTime(t *testing.T) {
	var ti time.Time
	fmt.Println(ti)
}
