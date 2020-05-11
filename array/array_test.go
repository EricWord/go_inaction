package array

import (
	"fmt"
	"testing"
)

/*
   @author eric
   @email 2635863525@qq.com
   @department SIOD DBA GROUP
   @date  2020/4/26 上午09:50
   @since Go1.13.8
*/

func TestArr(t *testing.T) {
	//声明一个包含5个元素的数组
	//var array [5]int

	//用具体的值初始化每个元素
	//arr:=[5]int{1,2,3,4,5}
	//容量由初始化的值的数量来决定
	//arr:=[...]int{1,2,3,4,5,6,6,7,8,9}
	//只是指定部分值，其余元素保持零值
	arr := [7]int{1: 9000, 2: 89999}
	fmt.Println(arr)

}

func TestArrayCopy(t *testing.T) {
	var arr1 [5]string
	arr2 := [5]string{"sdfvds", "sdfsdfv"}
	arr1 = arr2
	fmt.Println(arr1)

}
