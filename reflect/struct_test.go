package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

/*
   @author cuiguangsong
   @email cuiguangsong@duxiaoman.com
   @department SIOD DBA GROUP
   @date  2020/5/10 下午21:45
   @since Go1.13.8
*/

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func TestStruct(t *testing.T) {

	var m manager
	ty := reflect.TypeOf(&m)

	//获取指针的基类型
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	for i := 0; i < ty.NumField(); i++ {
		f := ty.Field(i)

		fmt.Println(f.Name, f.Type, f.Offset)

		//输出匿名字段结构
		if f.Anonymous {

			for x := 0; x < f.Type.NumField(); x++ {

				af := f.Type.Field(x)

				fmt.Println(" ", af.Name, af.Type)

			}

		}

	}

}
